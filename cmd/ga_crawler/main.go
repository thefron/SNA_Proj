package main

import (
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
