package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func baseName(year int, month int, day int, hour int) (name string, err error) {
	name = fmt.Sprintf("%d-%02d-%02d-%d", year, month, day, hour)
	layout := "2006-01-02-15"
	check, err := time.Parse(layout, name)
	if err != nil {
		return name, err
	}

	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		return name, err
	}

	max := time.Date(2016, 04, 30, 23, 1, 0, 0, utcLoc)
	min := time.Date(2015, 01, 01, 0, 0, 0, 0, utcLoc)
	if check.After(max) || check.Before(min) {
		return name, errors.New(fmt.Sprintf("%s is a not valid date.", name))
	}
	return name, nil
}

func jsonFileName(baseName string) string {
	return fmt.Sprintf("%s.json", baseName)
}

func gzipFileName(baseName string) string {
	jsonName := jsonFileName(baseName)
	return fmt.Sprintf("%s.gz", jsonName)
}

func getUrl(baseName string) (url string, err error) {
	gzName := gzipFileName(baseName)
	if err != nil {
		return url, err
	}
	return fmt.Sprintf("http://data.githubarchive.org/%s", gzName), nil
}

func download(dirPath string, baseName string) error {
	url, err := getUrl(baseName)
	if err != nil {
		return err
	}

	path := path.Join(dirPath, gzipFileName(baseName))
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	get, err := http.Get(url)
	if err != nil {
		return err
	}
	defer get.Body.Close()

	_, err = io.Copy(file, get.Body)
	return err
}

func unzip(dirPath string, baseName string) error {
	gzipName := gzipFileName(baseName)
	jsonName := jsonFileName(baseName)
	gzipPath := path.Join(dirPath, gzipName)
	jsonPath := path.Join(dirPath, jsonName)

	gzipFile, err := os.Open(gzipPath)
	if err != nil {
		return err
	}
	defer gzipFile.Close()

	jsonWriter, err := os.Create(jsonPath)
	if err != nil {
		return err
	}
	defer jsonWriter.Close()

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	_, err = io.Copy(jsonWriter, gzipReader)
	if err != nil {
		return err
	}

	return nil
}

func main() {
}

type User struct {
	Id    uint64 `json:"id"`
	Login string `json:"login"`
}
type Repo struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	Type      string    `json:"type"`
	Actor     User      `json:"actor"`
	Repo      Repo      `json:"repo"`
	CreatedAt time.Time `json:"created_at"`
	Org       *User     `json:"org"`
}

func readLines(filename string, c chan<- *Event) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for {
		event, err := readLine(scanner)
		if err != nil {
			return err
		}

		if event == nil {
			return nil
		}

		c <- event
	}
	return nil
}

func readLine(scanner *bufio.Scanner) (*Event, error) {
	if !scanner.Scan() {
		return nil, scanner.Err()
	}

	line := scanner.Bytes()

	event := Event{}
	if err := json.Unmarshal(line, &event); err != nil {
		return nil, err
	}

	return &event, nil
}

func writeEvents(filename string, c <-chan *Event) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	return func() error {
		for event := range c {
			err := writeEvent(writer, event)
			if err != nil {
				return err
			}
		}
		return nil
	}()
}

func writeEvent(writer *bufio.Writer, event *Event) error {
	createdAt := event.CreatedAt.Format("2006-01-02-15-04-05")
	content := fmt.Sprintf("%s %s %d %s %d %s", event.Type, event.Actor.Login, event.Actor.Id, event.Repo.Name, event.Repo.Id, createdAt)
	if event.Org != nil {
		content = fmt.Sprintf("%s %s %d\n", content, event.Org.Login, event.Org.Id)
	} else {
		content = fmt.Sprintf("%s\n", content)
	}
	_, err := writer.WriteString(content)
	return err
}
