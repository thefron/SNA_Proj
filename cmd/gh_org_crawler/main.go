package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Org struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type OrgWithMembers struct {
	Org
	Members []int
}

type User struct {
	Id    int
	Login string
}

func main() {
	var inputFileName string
	var outputFileName string
	var numberOfJobs int
	var token string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&outputFileName, "output", "output.txt", "output file name")
	flag.IntVar(&numberOfJobs, "jobs", 4, "number of jobs")
	flag.StringVar(&token, "token", "", "github token")
	flag.Parse()
	fmt.Println("Start with input:", inputFileName)
	fmt.Println("           output:", outputFileName)
	fmt.Println("           jobs:", numberOfJobs)
	fmt.Println("           token:", token)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	in := make(chan Org, numberOfJobs)
	out := make(chan OrgWithMembers)

	done := make(chan bool)
	go func(inputFile *os.File, in chan<- Org) {
		defer close(in)
		err := readLines(inputFile, in)
		if err != nil {
			fmt.Println(err)
			return
		}
	}(inputFile, in)
	go func(outputFile *os.File, out <-chan OrgWithMembers) {
		err := writeLines(outputFile, out)
		if err != nil {
			fmt.Println(err)
		}
		done <- true
	}(outputFile, out)

	for i := 0; i < numberOfJobs; i += 1 {
		go func(in <-chan Org, out chan<- OrgWithMembers, token string) {
			err = work(in, out, token)
			if err != nil {
				fmt.Println(err)
			}
		}(in, out, token)
	}
	<-done
}

func getIdsPage(orgName string, token string, page int) (*[]int, error) {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/public_members?page=%d", orgName, page)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	var resp *http.Response
	for {
		resp, err = client.Do(req)
		if resp.StatusCode == http.StatusOK {
			break
		}
		after := time.Duration(50 * time.Minute)
		resetTime, err := strconv.ParseInt(resp.Header.Get("X-Ratelimit-Reset"), 10, 64)
		if err == nil {
			after = time.Unix(resetTime, 0).Sub(time.Now())
		}
		fmt.Printf("Status = %d\n", resp.StatusCode)
		fmt.Println("Header", resp.Header)
		fmt.Println("Wait ", after)
		<-time.After(after)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	members := make([]User, 0)
	err = json.Unmarshal(body, &members)
	if err != nil {
		return nil, err
	}

	lengthOfMembers := len(members)
	memberIds := make([]int, lengthOfMembers)
	for i := 0; i < lengthOfMembers; i += 1 {
		memberIds[i] = members[i].Id
	}
	return &memberIds, nil
}

func getIds(orgName string, token string) (*[]int, error) {
	page := 1
	result := make([]int, 0)
	for {
		ids, err := getIdsPage(orgName, token, page)
		if err != nil {
			return nil, err
		}
		result = append(result, *ids...)
		if len(*ids) != 30 {
			break
		}
		page += 1
	}
	return &result, nil
}

func work(in <-chan Org, out chan<- OrgWithMembers, token string) error {
	defer close(out)
	for org := range in {
		ids, err := getIds(org.Name, token)
		if err != nil {
			return err
		}

		out <- OrgWithMembers{
			org,
			*ids,
		}
	}
	return nil
}

func readLines(inputFile *os.File, c chan<- Org) error {
	scanner := bufio.NewScanner(inputFile)
	for {
		if !scanner.Scan() {
			return scanner.Err()
		}
		line := scanner.Text()
		var org Org
		_, err := fmt.Sscanf(line, "%d %s", &org.Id, &org.Name)
		if err != nil {
			return err
		}
		c <- org
	}
}

func writeLines(outputFile *os.File, c <-chan OrgWithMembers) error {
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for org := range c {
		line := fmt.Sprintf("%d %s", org.Id, org.Name)
		for member := range org.Members {
			line = fmt.Sprintf("%s %d", line, org.Members[member])
		}
		line = fmt.Sprintf("%s\n", line)
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
