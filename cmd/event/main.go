package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func shouldNot(e error) {
	if e != nil {
		panic(e)
	}
}

type item struct {
	user int
	date string
}

func main() {
	var inputFileName string
	var outputFileName string
	var langFileName string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&langFileName, "lang", "lang.txt", "lang dir name")
	flag.StringVar(&outputFileName, "output", "output.txt", "output file name")
	flag.Parse()
	fmt.Println("Start with input:", inputFileName)
	fmt.Println("	         lang:", langFileName)
	fmt.Println("	       output:", outputFileName)

	inputFile, err := os.Open(inputFileName)
	shouldNot(err)
	defer inputFile.Close()

	langFile, err := os.Open(langFileName)
	shouldNot(err)
	defer langFile.Close()
	reader := bufio.NewReader(langFile)
	scanner := bufio.NewScanner(reader)
	langs := make(map[int]string, 0)
	for scanner.Scan() {
		shouldNot(scanner.Err())
		tokens := strings.Split(scanner.Text(), " ")
		repoId, err := strconv.Atoi(tokens[0])
		shouldNot(err)
		langName := tokens[2]
		langs[repoId] = langName
	}

	weakHandled := make(map[string](map[int]int), 0)
	weakList := make(map[string]([]item), 0)
	strongHandled := make(map[string](map[int]int), 0)
	strongList := make(map[string]([]item), 0)

	scanner = bufio.NewScanner(inputFile)
	for scanner.Scan() {
		func() {
			defer func() {
				r := recover()
				if r != nil {
					fmt.Println(r)
				}
			}()
			shouldNot(scanner.Err())
			line := scanner.Text()
			tokens := strings.Split(line, " ")
			eventName := tokens[0]
			userId, err := strconv.Atoi(tokens[2])
			shouldNot(err)
			repoId, err := strconv.Atoi(tokens[4])
			shouldNot(err)
			date := tokens[5]
			langName, ok := langs[repoId]
			if !ok {
				return
			}
			handled, ok := weakHandled[langName]
			if !ok {
				handled = make(map[int]int, 0)
			}
			list, ok := weakList[langName]
			if !ok {
				list = []item{}
			}

			count, ok := handled[userId]
			if ok {
				handled[userId] = count + 1
			} else {
				handled[userId] = 1
				list = append(list, item{userId, date})
			}
			weakHandled[langName] = handled
			weakList[langName] = list

			if !isStrong(eventName) {
				return
			}

			handled, ok = strongHandled[langName]
			if !ok {
				handled = make(map[int]int, 0)
			}
			list, ok = strongList[langName]
			if !ok {
				list = []item{}
			}

			count, ok = handled[userId]
			if ok {
				handled[userId] = count + 1
			} else {
				handled[userId] = 1
				list = append(list, item{userId, date})
			}
			strongHandled[langName] = handled
			strongList[langName] = list
		}()
	}
	fmt.Println("Parsed")
	// fmt.Println("weak handle count\n")

	// for lang, handled := range weakHandled {
	// 	count1 := 0
	// 	count10 := 0
	// 	count100 := 0
	// 	count1000 := 0
	// 	count10000 := 0
	// 	for _, count := range handled {
	// 		count1 += 1
	// 		if count >= 10 {
	// 			count10 += 1
	// 		}
	// 		if count >= 100 {
	// 			count100 += 1
	// 		}
	// 		if count >= 1000 {
	// 			count1000 += 1
	// 		}
	// 		if count >= 10000 {
	// 			count10000 += 1
	// 		}
	// 	}
	// 	fmt.Println(lang, "number of weak :", count1, count10, count100, count1000, count10000)
	// }
	// fmt.Println("strong handle count\n")
	// for lang, handled := range strongHandled {
	// 	count1 := 0
	// 	count10 := 0
	// 	count100 := 0
	// 	count1000 := 0
	// 	count10000 := 0
	// 	for _, count := range handled {
	// 		count1 += 1
	// 		if count >= 10 {
	// 			count10 += 1
	// 		}
	// 		if count >= 100 {
	// 			count100 += 1
	// 		}
	// 		if count >= 1000 {
	// 			count1000 += 1
	// 		}
	// 		if count >= 10000 {
	// 			count10000 += 1
	// 		}
	// 	}
	// 	fmt.Println(lang, "number of strong :", count1, count10, count100, count1000, count10000)
	// }

	for lang, list := range weakList {
		fmt.Println("weak ", lang)
		func() {
			defer func() {
				r := recover()
				if r != nil {
					fmt.Println(r)
				}
			}()
			counter, ok := weakHandled[lang]
			if !ok {
				return
			}
			fileName := fmt.Sprintf("weak/%s.txt", strings.Replace(strings.Replace(lang, "/", "-", -1), " ", "-", -1))
			file, err := os.Create(fileName)
			defer file.Close()
			shouldNot(err)
			writer := bufio.NewWriter(file)
			for _, item := range list {
				func() {
					defer func() {
						r := recover()
						if r != nil {
							fmt.Println(r)
						}
					}()
					count, ok := counter[item.user]
					if !ok {
						return
					}
					if count < 10 {
						return
					}
					writer.WriteString(fmt.Sprintf("%d %s %d\n", item.user, item.date, count))
					writer.Flush()
				}()
			}
		}()
	}

	for lang, list := range strongList {
		fmt.Println("strong ", lang)
		func() {
			defer func() {
				r := recover()
				if r != nil {
					fmt.Println(r)
				}
			}()
			counter, ok := strongHandled[lang]
			if !ok {
				return
			}
			fileName := fmt.Sprintf("strong/%s.txt", strings.Replace(strings.Replace(lang, "/", "-", -1), " ", "-", -1))
			file, err := os.Create(fileName)
			defer file.Close()
			shouldNot(err)
			writer := bufio.NewWriter(file)
			for _, item := range list {
				func() {
					defer func() {
						r := recover()
						if r != nil {
							fmt.Println(r)
						}
					}()
					count, ok := counter[item.user]
					if !ok {
						return
					}
					if count < 10 {
						return
					}
					writer.WriteString(fmt.Sprintf("%d %s %d\n", item.user, item.date, count))
					writer.Flush()
				}()
			}
		}()
	}
}

func isStrong(eventName string) bool {
	switch eventName {
	case "PullRequestEvent":
	case "PushEvent":
		return true
	case "DeploymentEvent":
	case "DeploymentStatusEvent":
	case "ForkEvent":
	case "WatchEvent":
	case "CreateEvent":
	case "DeleteEvent":
	case "GollumEvent":
	case "IssuesEvent":
	case "IssueCommentEvent":
	case "CommitCommentEvent":
	case "MemberEvent":
	case "MembershipEvent":
	case "PageBuildEvent":
	case "PublicEvent":
	case "PullRequestReviewCommentEvent":
	case "ReleaseEvent":
	case "RepositoryEvent":
	case "StatusEvent":
	case "TeamAddEvent":
		return false
	default:
		panic(errors.New(fmt.Sprintf("%s is not handled", eventName)))
	}
	return false
}
