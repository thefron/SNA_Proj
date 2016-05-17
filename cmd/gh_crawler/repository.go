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

func main() {
	var inputFileName string
	var tokenFileName string
	var languageFileName string
	var forksFileName string
	var stargazersFileName string
	var subscribersFileName string
	flag.StringVar(&inputFileName, "input", "input-repo.txt", "input file (repository) name")
	flag.StringVar(&tokenFileName, "token", "tokens.txt", "files storing github tokens")
	flag.StringVar(&languageFileName, "language", "language.txt", "output language file")
	flag.StringVar(&forksFileName, "forks", "forks.txt", "output forks file")
	flag.StringVar(&stargazersFileName, "stargazers", "stargazers.txt", "output stargazers file")
	flag.StringVar(&subscribersFileName, "subscribers", "subscribers.txt", "output subscribers file")
	flag.Parse()
	fmt.Println("repository crawler")
	fmt.Println("input file:", inputFileName)
	fmt.Println("token file:", tokenFileName)
	fmt.Println("language output:", languageFileName)
	fmt.Println("forks output:", forksFileName)
	fmt.Println("stargazers output:", stargazersFileName)
	fmt.Println("subscribers output:", subscribersFileName)

	tokens, err := readTokens(tokenFileName)
	if err != nil {
		panic(err)
	}

	numberOfTokens := len(tokens)
	fmt.Println("total tokens:", numberOfTokens)

	readerCh := make(chan *Repo)
	writerCh := make(chan *RepoInfo)

	for i := 0; i < numberOfTokens; i++ {
		go fetchRepositoryInfo(readerCh, writerCh, tokens[i])
	}

	go writeResult(writerCh, languageFileName, forksFileName, stargazersFileName, subscribersFileName)

	readRepoLines(readerCh, inputFileName)

	wg.Wait()
	log.Println("Done")
}

func writeResult(channel <-chan *RepoInfo, languageFileName string, forksFileName string, stargazersFileName string, subscribersFileName string) {
	languageFile, err := os.Create(languageFileName)
	if err != nil {
		panic(err)
	}
	defer languageFile.Close()

	forksFile, err := os.Create(forksFileName)
	if err != nil {
		panic(err)
	}
	defer forksFile.Close()

	stargazersFile, err := os.Create(stargazersFileName)
	if err != nil {
		panic(err)
	}
	defer stargazersFile.Close()

	subscribersFile, err := os.Create(subscribersFileName)
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

	wg.Add(1)

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

func needRetry(result *octokit.Result) bool {
	if !result.HasError() {
		return false
	}
	if result.Response == nil {
		return false
	}
	if result.RateLimitRemaining() == 0 {
		log.Println("rate limit reached")
		resetTime := result.RateLimitReset()
		after := resetTime.Sub(time.Now()) + (10 * time.Second)
		log.Println("waiting", after)
		<-time.After(after)
		return true
	}
	return false
}

func fetchAllUserID(client *octokit.Client, url *url.URL) []int {
	var allID []int
	for {
		users, result := client.Users(url).All()
		if needRetry(result) {
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
		repositories, result := client.Repositories().All(link, params)
		if needRetry(result) {
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



func fetchRepositoryInfo(repoCh <-chan *Repo, resultCh chan<- *RepoInfo, tokenString string) {
	defer close(resultCh)

	var (
		repoURL = octokit.Hyperlink("repos/{owner}/{repo}") /* repository */
		forksURL = octokit.Hyperlink("repos/{owner}/{repo}/forks") /* repositories */
		stargazersURL = octokit.Hyperlink("repos/{owner}/{repo}/stargazers") /* users */
		subscribersURL = octokit.Hyperlink("repos/{owner}/{repo}/subscribers") /* users */
		//contributorsURL = octokit.Hyperlink("repos/{owner}/{repo}/contributors") /* users */
	)

	token := octokit.TokenAuth { tokenString }
	client := octokit.NewClient(token)

	for repo := range repoCh {
		fmt.Println("processing ", repo.Name)
		param := octokit.M{"owner": repo.Owner(), "repo": repo.Repo() }

		var generalRepoInfo *octokit.Repository
		var result *octokit.Result

		for {
			generalRepoInfo, result = client.Repositories().One(&repoURL, param)
			if needRetry(result) {
				continue
			}
			break
		}

		if generalRepoInfo == nil {
			fmt.Println("repository not found", repo.Name)
			continue
		}

		var url *url.URL
		var err error

		url, err = stargazersURL.Expand(param)
		if err != nil {
			panic(err)
		}
		stargazers := fetchAllUserID(client, url)

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
		repoInfo.Stargazers = stargazers
		repoInfo.Subscribers = subscribers
		repoInfo.Contributors = nil

		resultCh <- repoInfo
	}
}

