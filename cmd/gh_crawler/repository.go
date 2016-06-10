package main

import (
	"sync"
	"time"
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
	"flag"
	"net/url"
	"net/http"
	"bytes"
	"github.com/octokit/go-octokit/octokit"
)

type Repo struct {
	ID int
	Name string
}

type RepoInfo struct {
	ID int
	Name string
	Language string
	Forkers []int
	Stargazers []int
	Subscribers []int
	Contributors []int
}

func (r *Repo) Owner() string {
	splits := strings.Split(r.Name, "/")
	return splits[0]
}

func (r *Repo) Repo() string {
	splits := strings.Split(r.Name, "/")
	return splits[1]
}

func readTokens(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tokens := make([]string, 0)
	for {
		if !scanner.Scan() {
			return tokens, scanner.Err()
		}
		token := strings.Trim(scanner.Text(), " \r\n\t")
		tokens = append(tokens, token)
	}
}

var wg sync.WaitGroup
var wg_tokens sync.WaitGroup

func main() {
	var inputFileName string
	var tokenFileName string
	var languageFileName string
	var forksFileName string
	var stargazersFileName string
	var subscribersFileName string
	var progressFileName string
	flag.StringVar(&inputFileName, "input", "input-repo.txt", "input file (repository) name")
	flag.StringVar(&tokenFileName, "token", "tokens.txt", "files storing github tokens")
	flag.StringVar(&languageFileName, "language", "language.txt", "output language file")
	flag.StringVar(&forksFileName, "forks", "forks.txt", "output forks file")
	flag.StringVar(&stargazersFileName, "stargazers", "stargazers.txt", "output stargazers file")
	flag.StringVar(&subscribersFileName, "subscribers", "subscribers.txt", "output subscribers file")
	flag.StringVar(&progressFileName, "progress", "progress.txt", "saving current progress")
	flag.Parse()
	fmt.Println("repository crawler")
	fmt.Println("input file:", inputFileName)
	fmt.Println("token file:", tokenFileName)
	fmt.Println("language output:", languageFileName)
	fmt.Println("forks output:", forksFileName)
	fmt.Println("stargazers output:", stargazersFileName)
	fmt.Println("subscribers output:", subscribersFileName)
	fmt.Println("progress file:", progressFileName)

	tokens, err := readTokens(tokenFileName)
	if err != nil {
		panic(err)
	}

	numberOfTokens := len(tokens)
	fmt.Println("total tokens:", numberOfTokens)

	readerCh := make(chan *Repo)
	writerCh := make(chan *RepoInfo)
	progressCh := make(chan int)

	doneSet := getProgress(progressFileName)

	for i := 0; i < numberOfTokens; i++ {
		wg_tokens.Add(1)
		go fetchRepositoryInfo(readerCh, writerCh, tokens[i], doneSet, progressCh)
	}

	go writeResult(writerCh, languageFileName, forksFileName, stargazersFileName, subscribersFileName)

	go writeProgress(progressCh, progressFileName)

	readRepoLines(readerCh, inputFileName)

	/* works of fetchRepositories done */
	wg_tokens.Wait()
	/* close writing channels */
	close(progressCh)
	close(writerCh)

	/* all writing tasks done */
	wg.Wait()
	log.Println("Done")
}

func openAppend(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
}

func writeProgress(channel <-chan int, progressFileName string) {
	progressFile, err := openAppend(progressFileName)
	if err != nil {
		panic(err)
	}
	defer progressFile.Close()

	for id := range channel {
		progressFile.WriteString(fmt.Sprintf("%d\n", id))
	}
	wg.Done()
}

func writeResult(channel <-chan *RepoInfo, languageFileName string, forksFileName string, stargazersFileName string, subscribersFileName string) {
	languageFile, err := openAppend(languageFileName)
	if err != nil {
		panic(err)
	}
	defer languageFile.Close()

	forksFile, err := openAppend(forksFileName)
	if err != nil {
		panic(err)
	}
	defer forksFile.Close()

	stargazersFile, err := openAppend(stargazersFileName)
	if err != nil {
		panic(err)
	}
	defer stargazersFile.Close()

	subscribersFile, err := openAppend(subscribersFileName)
	if err != nil {
		panic(err)
	}
	defer subscribersFile.Close()


	for repoInfo := range channel {
		if repoInfo.Language != "" {
			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("%d %s", repoInfo.ID, repoInfo.Name))
			buffer.WriteString(fmt.Sprintf(" %s\n", repoInfo.Language))
			_, err := buffer.WriteTo(languageFile)
			if err != nil {
				log.Println(err)
			}
		}

		idPrinter := func(idList []int, outputFile *os.File) {
			if idList == nil {
				return
			}
			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("%d %s", repoInfo.ID, repoInfo.Name))
			for _, id := range idList {
				buffer.WriteString(fmt.Sprintf(" %d", id))
			}
			buffer.WriteString("\n")
			_, err := buffer.WriteTo(outputFile)
			if err != nil {
				log.Println(err)
			}
		}

		idPrinter(repoInfo.Forkers, forksFile)
		idPrinter(repoInfo.Stargazers, stargazersFile)
		idPrinter(repoInfo.Subscribers, subscribersFile)
	}
	wg.Done()
}

func readRepoLines(channel chan<- *Repo, inputFileName string) error {
	defer close(channel)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(make([]byte, 1024), bufio.MaxScanTokenSize*10)

	for {
		repo, err := parseRepoLine(scanner)

		if err != nil {
			return err
		}

		if repo == nil {
			break
		}

		channel <- repo
	}

	wg.Add(2)

	return nil
}

func removeEmptyStrings(strs []string) []string {
	var result []string
	for _, str := range strs {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func parseRepoLine(scanner *bufio.Scanner) (*Repo, error) {
	for {
		if !scanner.Scan() {
			return nil, scanner.Err()
		}

		line := scanner.Text()
		splits := removeEmptyStrings(strings.Split(line, " "))

		// skip ill-formed line
		if len(splits) != 2 {
			continue
		}

		id, err := strconv.Atoi(splits[0])

		if err != nil {
			return nil, err
		}

		repo := &Repo{ID: id, Name: splits[1]}
		return repo, nil
	}
}

func getRateLimitResetTime(response *http.Response) *time.Time {
	epoc := response.Header.Get("X-RateLimit-Reset")
	if epoc == "" {
		return nil
	}
	reset, err := strconv.ParseInt(epoc, 10, 64)
	if err != nil {
		return nil
	}
	t := time.Unix(reset, 0)
	return &t
}

func needRetry(result *octokit.Result, token string) bool {
	if !result.HasError() {
		rate := result.RateLimitRemaining()
		if rate < 500+360 {
			after := 10 * time.Second
			<-time.After(after)
		}
		return false
	}
	rerr, ok := result.Err.(*octokit.ResponseError)
	if ok && rerr.Type == octokit.ErrorTooManyRequests {
		log.Println("rate limit reached")
		log.Println(result, token)
		resetTime := getRateLimitResetTime(rerr.Response)
		if resetTime == nil {
			log.Println("unknown reset time")
			after := (5 * time.Minute)
			log.Println("waiting", after)
			<-time.After(after)
		} else {
			after := resetTime.Sub(time.Now()) + (10 * time.Second)
			log.Println("waiting", after)
			<-time.After(after)
		}
		return true
	}
	// unauthorized -> means bad token
	if ok && rerr.Type == octokit.ErrorUnauthorized {
		log.Println("bad token!", token)
		panic(result)
	}
	if result.Response == nil {
		return false
	}
	return false
}

func fetchAllUserID(client *octokit.Client, url *url.URL) []int {
	var allID []int
	for {
		fmt.Println("query start");
		users, result := client.Users(url).All()
		if needRetry(result, client.AuthMethod.String()) {
			continue
		}
		if result.HasError() {
			return nil
		}

		for _, user := range users {
			allID = append(allID, user.ID)
		}

		if result.NextPage == nil {
			break
		}
		url, _ = result.NextPage.Expand(nil)
	}
	return allID
}

func fetchAllRepositoriesOwners(client *octokit.Client, link *octokit.Hyperlink, params octokit.M) []int {
	var allID []int
	for {
		fmt.Println("query start");
		repositories, result := client.Repositories().All(link, params)
		if needRetry(result, client.AuthMethod.String()) {
			continue
		}
		if result.HasError() {
			return nil
		}

		for _, repository := range repositories {
			allID = append(allID, repository.Owner.ID)
		}

		if result.NextPage == nil {
			break
		}
		link = result.NextPage
	}
	return allID
}


func getProgress(progressFileName string) map[int]bool {
	doneSet := map[int]bool{}
	progressFile, err := os.Open(progressFileName)
	if err != nil {
		return doneSet
	}
	defer progressFile.Close()

	scanner := bufio.NewScanner(progressFile)
	scanner.Buffer(make([]byte, 1024), bufio.MaxScanTokenSize*10)
	for {
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		id, err := strconv.Atoi(line)

		if err != nil {
			continue
		}

		doneSet[id] = true
	}

	log.Println("Found:", len(doneSet), "records")

	return doneSet
}

func fetchRepositoryInfo(repoCh <-chan *Repo, resultCh chan<- *RepoInfo, tokenString string, done map[int]bool, progCh chan<- int) {
	defer wg_tokens.Done()

	var (
		repoURL = octokit.Hyperlink("repos/{owner}/{repo}?access_token={token}") /* repository */
		forksURL = octokit.Hyperlink("repos/{owner}/{repo}/forks?access_token={token}") /* repositories */
		//stargazersURL = octokit.Hyperlink("repos/{owner}/{repo}/stargazers") /* users */
		subscribersURL = octokit.Hyperlink("repos/{owner}/{repo}/subscribers?access_token={token}") /* users */
		//contributorsURL = octokit.Hyperlink("repos/{owner}/{repo}/contributors") /* users */
	)

	token := octokit.TokenAuth { tokenString }
	client := octokit.NewClient(token)

	for repo := range repoCh {
		_, doneBefore := done[repo.ID]
		if doneBefore {
			continue
		}
		log.Println("processing ", repo.Name)
		param := octokit.M{"owner": repo.Owner(), "repo": repo.Repo(), "token": tokenString }

		var generalRepoInfo *octokit.Repository
		var result *octokit.Result

		for {
			fmt.Println("query start");
			generalRepoInfo, result = client.Repositories().One(&repoURL, param)
			if needRetry(result, client.AuthMethod.String()) {
				continue
			}
			break
		}

		if generalRepoInfo == nil {
			log.Println("repository not found", repo.Name)
			progCh<- repo.ID
			continue
		}

		var url *url.URL
		var err error

		/*url, err = stargazersURL.Expand(param)
		if err != nil {
			panic(err)
		}
		stargazers := fetchAllUserID(client, url) */

		url, err = subscribersURL.Expand(param)
		if err != nil {
			panic(err)
		}
		subscribers := fetchAllUserID(client, url)

		forkers := fetchAllRepositoriesOwners(client, &forksURL, param)

		repoInfo := new(RepoInfo)
		repoInfo.ID = repo.ID
		repoInfo.Name = repo.Name
		repoInfo.Language = generalRepoInfo.Language
		repoInfo.Forkers = forkers
		repoInfo.Stargazers = nil //stargazers
		repoInfo.Subscribers = subscribers
		repoInfo.Contributors = nil

		resultCh <- repoInfo
		progCh<- repo.ID
	}
}

