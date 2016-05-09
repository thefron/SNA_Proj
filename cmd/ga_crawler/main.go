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
		downAndWriteEventsForDay(targetDay, downDir)
	}
}

func downAndWriteEventsForDay(date time.Time, downDir string) {
	c := make(chan *Event, 100)
	layout := "2006-01-02"
	fileName := path.Join(downDir, date.Format(layout))
	go func() {
		err := writeEvents(fileName, c)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in downAndWriteEventsForDay(%s) %s!\n", date.Format(layout), err)
		}
	}()
	repeat := 4
	for i := 0; i < 24; i += repeat {
		collectNHour(date, i, repeat, downDir, c)
	}
	close(c)
}

func collectNHour(time time.Time, start int, repeat int, downDir string, c chan<- *Event) {
	year := time.Year()
	month := time.Month()
	day := time.Day()
	done := make(chan int, 1)
	completeJob := 0
	for hour := start; hour < start+repeat; hour += 1 {
		go func(hour int) {
			err := downAndReadEvents(year, month, day, hour, downDir, c)
			completeJob += 1
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error in downAndReadEvents(%d-%d-%d-%d) %s!\n", year, month, day, hour, err)
			}
			if completeJob == repeat {
				done <- 1
			}
		}(hour)
	}
	<-done
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
