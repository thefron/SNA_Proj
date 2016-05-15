package main

import (
	"io/ioutil"
	"sync"
	"bytes"
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
	"strconv"
	"github.com/octokit/go-octokit/octokit"
)

type Repo struct {
	ID int
	Name string
}

func (r *Repo) Owner() string {
	splits := strings.Split(r.Name, "/")
	return splits[0]
}

func (r *Repo) Repo() string {
	splits := strings.Split(r.Name, "/")
	return splits[1]
}


func getAuthToken() octokit.TokenAuth {
	dat, err := ioutil.ReadFile("token.txt")
	if err != nil {
		fmt.Println("token.txt not found. create the file. go to https://github.com/settings/tokens -> put token string in the file")
		panic(err)
	}

	tokenString := strings.Trim(string(dat), " \r\n\t")

	return octokit.TokenAuth { tokenString }
}

var wg sync.WaitGroup

func main() {
	var languageFileName string = "languages.txt"
	var inputFileName string = "input-repo.txt"

	languagesRepoCh := make(chan *Repo)

	token := getAuthToken()
	client := octokit.NewClient(token)
	repoService := client.Repositories()

	languagesOutputFile, err := os.Create(languageFileName)
	if err != nil {
		panic(err)
	}
	defer languagesOutputFile.Close()

	go fetchLanguages(languagesOutputFile, languagesRepoCh, repoService)


	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	readRepoLines(inputFile, []chan<- *Repo {
		languagesRepoCh,
	})

	wg.Wait()
	log.Println("Done")


	//repo, _ := repoService.One(nil, octokit.M{"owner": "thefron", "repo": "SNA_Proj"})
	//fmt.Println(repo.Language)

}

func readRepoLines(inputFile *os.File, repoChs []chan<- *Repo) error {
	scanner := bufio.NewScanner(inputFile)
	scanner.Buffer(make([]byte, 1024), bufio.MaxScanTokenSize*10)

	defer func() {
		for _, ch := range repoChs {
			close(ch)
		}
	}()

	for {
		repo, err := parseRepoLine(scanner)

		if err != nil {
			return err
		}

		if repo == nil {
			break
		}

		for _, ch := range repoChs {
			wg.Add(1)
			ch <- repo
		}
	}

	return nil
}

func parseRepoLine(scanner *bufio.Scanner) (*Repo, error) {
	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	line := scanner.Text()
	splits := strings.Split(line, " ")

	id, err := strconv.Atoi(splits[0])

	if err != nil {
		return nil, err
	}

	repo := &Repo{ID: id, Name: splits[1]}
	return repo, nil
}


func fetchLanguages(outputFile *os.File, repoCh <-chan *Repo, repoService *octokit.RepositoriesService) {
	for repo := range repoCh {
		go func(repo *Repo) {
			repoInfo, _ := repoService.One(nil, octokit.M{"owner": repo.Owner(), "repo": repo.Repo() })
			language := repoInfo.Language
			var buffer bytes.Buffer
			buffer.WriteString(fmt.Sprintf("%d %s", repo.ID, repo.Name))
			buffer.WriteString(fmt.Sprintf(" %s", language))
			buffer.WriteString("\n")
			_, err := buffer.WriteTo(outputFile)
			if err != nil {
				log.Println(err)
			}

			wg.Done()
		}(repo)
	}
}
