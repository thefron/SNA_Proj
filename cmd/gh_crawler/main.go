package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/octokit/go-octokit/octokit"
)

type User struct {
	ID    int
	Login string
}

var wg sync.WaitGroup
var client = octokit.NewClient(nil)
var followersService = client.Followers()

func main() {
	var inputFileName string
	var followersOutputFileName string
	var followingsOutputFileName string

	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&followersOutputFileName, "followers-output", "followers.txt", "followers output file name")
	flag.StringVar(&followingsOutputFileName, "followings-output", "followings.txt", "followings output file name")

	followersUserCh := make(chan *User)
	followingsUserCh := make(chan *User)

	followersOutputFile, err := os.Create("followers.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer followersOutputFile.Close()

	go fetchFollowers(followersOutputFile, followersUserCh)

	followingsOutputFile, err := os.Create("followings.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer followingsOutputFile.Close()

	go fetchFollowings(followingsOutputFile, followingsUserCh)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()

	readUserLines(inputFile, []chan<- *User{
		followersUserCh,
		followingsUserCh,
	})

	wg.Wait()
	log.Println("Done")
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

func fetchFollowers(outputFile *os.File, userCh <-chan *User) {
	for user := range userCh {
		go func(user *User) {
			followers, _ := followersService.All(&octokit.FollowerUrl, octokit.M{"user": user.Login})

			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("%d %s", user.ID, user.Login))
			for _, follower := range followers {
				buffer.WriteString(" " + strconv.Itoa(follower.ID))
			}
			buffer.WriteString("\n")
			_, err := buffer.WriteTo(outputFile)
			if err != nil {
				log.Println(err)
			}

			wg.Done()
		}(user)
	}
}

func fetchFollowings(outputFile *os.File, userCh <-chan *User) {
	for user := range userCh {
		go func(user *User) {
			followings, _ := followersService.All(&octokit.FollowingUrl, octokit.M{"user": user.Login})

			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("%d %s", user.ID, user.Login))
			for _, follower := range followings {
				buffer.WriteString(" " + strconv.Itoa(follower.ID))
			}
			buffer.WriteString("\n")
			_, err := buffer.WriteTo(outputFile)
			if err != nil {
				log.Println(err)
			}

			wg.Done()
		}(user)
	}
}
