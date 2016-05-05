package main

import (
	"bufio"
	"fmt"
	"os"
)

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
