package main

import (
	"errors"
	"fmt"
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

func main() {
}
