package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

func main() {
	if len(os.Args) != 4 {
		panic("Usage: {HH-MM-DD: start day} {HH-MM-DD: end day} {download directory}")
	}

	layout := "2006-01-02"

	start, err := time.Parse(layout, os.Args[1])
	if err != nil {
		panic("The first argument should be start day")
	}

	end, err := time.Parse(layout, os.Args[2])
	if err != nil {
		panic("The second argument should be end day")
	}

	downDir := os.Args[3]

	for targetDay := end.AddDate(0, 0, -1); start.Unix() <= targetDay.Unix(); targetDay = targetDay.AddDate(0, 0, -1) {
		err = downAndWriteEventsForDay(targetDay, downDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in downAndWriteEventsForDay(%s) %s!\n", targetDay.Format(layout), err)
		}
	}
}

func downAndWriteEventsForDay(date time.Time, downDir string) error {
	c := make(chan *Event, 100)
	year := date.Year()
	month := date.Month()
	day := date.Day()
	completeJob := 0
	fileName := path.Join(downDir, fmt.Sprintf("%d-%02d-%02d", year, month, day))
	for hour := 0; hour < 24; hour += 1 {
		go func(hour int) {
			err := downAndReadEvents(year, month, day, hour, downDir, c)
			completeJob += 1
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in downAndReadEvents(%d-%d-%d-%d) %s!\n", year, month, day, hour, err)
			}
			if completeJob == 24 {
				close(c)
			}
		}(hour)
	}
	return writeEvents(fileName, c)
}

func downAndReadEvents(year int, month time.Month, day int, hour int, downDir string, c chan<- *Event) error {
	targetName, err := baseName(year, month, day, hour)
	if err != nil {
		return err
	}
	err = download(downDir, targetName)
	if err != nil {
		return err
	}
	jsonFile := path.Join(downDir, jsonFileName(targetName))
	gzipFile := path.Join(downDir, gzipFileName(targetName))
	unzip(downDir, targetName)
	os.Remove(gzipFile)
	defer os.Remove(jsonFile)

	return readLines(jsonFile, c)
}
