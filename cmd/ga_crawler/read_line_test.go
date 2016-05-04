package main

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReadLine(t *testing.T) {
	file, err := os.Open("./test_data/one_line_sample.json")
	assert.Nil(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	event, err := readLine(scanner)
	assert.Nil(t, err)
	assert.NotNil(t, event)

	assert.Equal(t, "IssuesEvent", event.Type)
	assert.Equal(t, uint64(793041), event.Actor.Id)
	assert.Equal(t, "OndraM", event.Actor.Login)
	assert.Equal(t, uint64(35304655), event.Repo.Id)
	assert.Equal(t, "lmc-eu/steward", event.Repo.Name)
	utc, err := time.LoadLocation("UTC")
	assert.Nil(t, err)
	assert.Equal(t, time.Date(2016, 1, 1, 23, 0, 0, 0, utc), event.CreatedAt)
	assert.Equal(t, uint64(429615), event.Org.Id)
	assert.Equal(t, "lmc-eu", event.Org.Login)

	event, err = readLine(scanner)
	assert.Nil(t, err)
	assert.Nil(t, event)
}
