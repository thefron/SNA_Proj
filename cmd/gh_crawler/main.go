package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/octokit/go-octokit/octokit"
)

type User struct {
	ID    int
	Login string
}

type UserRelation struct {
	user      *User
	following []int
	followers []int
}

var wg sync.WaitGroup

func main() {
	var inputFileName string
	var followersOutputFileName string
	var followingOutputFileName string
	var tokensFileName string

	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&followersOutputFileName, "followers-output", "followers.txt", "followers output file name")
	flag.StringVar(&followingOutputFileName, "following-output", "following.txt", "following output file name")
	flag.StringVar(&tokensFileName, "tokens-file", "tokens.txt", "tokens file name")

	tokens, err := readTokens(tokensFileName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("total tokens:", len(tokens))

	userCh := make(chan *User)
	userRelationCh := make(chan *UserRelation)

	for idx, token := range tokens {
		log.Printf("Spawning jobs #%d", idx)
		go fetchRelations(token, userCh, userRelationCh, idx)
	}

	followersOutputFile, err := os.Create(followersOutputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer followersOutputFile.Close()

	followingOutputFile, err := os.Create(followingOutputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer followingOutputFile.Close()

	go writeUserRelations(followersOutputFile, followingOutputFile, userRelationCh)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()

	readUserLines(inputFile, []chan<- *User{userCh})

	wg.Wait()
	log.Println("Done")
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
		token := strings.Trim(scanner.Text(), "\n")
		tokens = append(tokens, token)
	}
}

func readUserLines(inputFile *os.File, userChs []chan<- *User) error {
	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(make([]byte, 1024), bufio.MaxScanTokenSize*10)

	for {
		user, err := parseUserLine(scanner)

		if user == nil {
			break
		}

		if err != nil {
			return err
		}

		if user == nil {
			log.Println("nil User")
		}

		for _, ch := range userChs {
			wg.Add(1)
			ch <- user
		}
	}

	for _, ch := range userChs {
		close(ch)
	}

	return nil
}

func parseUserLine(scanner *bufio.Scanner) (*User, error) {
	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	line := scanner.Text()
	splits := strings.Split(line, " ")

	id, err := strconv.Atoi(splits[0])

	if err != nil {
		return nil, err
	}

	user := &User{ID: id, Login: splits[1]}
	return user, nil
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

func fetchRelations(tokenString string, userCh <-chan *User, userRelationCh chan<- *UserRelation, jobIdx int) {

	var (
		followingUrl = octokit.Hyperlink("users/{user}/following?access_token={token}")
		followersUrl = octokit.Hyperlink("users/{user}/followers?access_token={token}")
	)

	token := octokit.TokenAuth{tokenString}
	client := octokit.NewClient(token)

	for user := range userCh {
		followingIDs := fetchFollowersForUser(client, &followingUrl, octokit.M{"user": user.Login, "token": tokenString})
		followersIDs := fetchFollowersForUser(client, &followersUrl, octokit.M{"user": user.Login, "token": tokenString})

		userRelation := &UserRelation{
			user:      user,
			following: followingIDs,
			followers: followersIDs,
		}

		userRelationCh <- userRelation
	}
}

func fetchFollowersForUser(client *octokit.Client, link *octokit.Hyperlink, params octokit.M) []int {
	var followerIDs []int
	for {
		followers, result := client.Followers().All(link, params)
		if needRetry(result, client.AuthMethod.String()) {
			continue
		}

		if result.HasError() {
			return nil
		}

		for _, user := range followers {
			followerIDs = append(followerIDs, user.ID)
		}

		if result.NextPage == nil {
			break
		}

		link = result.NextPage
	}

	return followerIDs
}

func writeUserRelations(followersOutputFile *os.File, followingOutputFile *os.File, userRelationCh <-chan *UserRelation) {
	for userRelation := range userRelationCh {
		writeRelation(followersOutputFile, userRelation.user, userRelation.followers)
		writeRelation(followingOutputFile, userRelation.user, userRelation.following)

		wg.Done()
	}
}

func writeRelation(outputFile *os.File, user *User, relatedUserIDs []int) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%d %s", user.ID, user.Login))
	for _, userID := range relatedUserIDs {
		buffer.WriteString(" " + strconv.Itoa(userID))
	}
	buffer.WriteString("\n")
	_, err := buffer.WriteTo(outputFile)
	if err != nil {
		log.Println(err)
	}
}
