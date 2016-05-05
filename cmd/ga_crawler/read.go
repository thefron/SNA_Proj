package main

import (
	"bufio"
	"encoding/json"
	"os"
)

func readLines(filename string, c chan<- *Event) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024), bufio.MaxScanTokenSize*10)

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
