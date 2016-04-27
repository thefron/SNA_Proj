package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	name, err := baseName(2016, 4, 20, 3)
	assert.Nil(t, err)
	assert.Equal(t, "2016-04-20-3", name)

	testDir := "./test_data"
	err = download(testDir, name)

	defer os.Remove(path.Join(testDir, gzipFileName(name)))
	assert.Nil(t, err)
}
