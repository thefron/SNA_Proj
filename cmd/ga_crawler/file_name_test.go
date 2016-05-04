package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseNameReturnsString(t *testing.T) {
	name, _ := baseName(2016, 1, 1, 3)
	assert.Equal(t, "2016-01-01-3", name)
}

func Test2016043023IsValid(t *testing.T) {
	name, err := baseName(2016, 4, 30, 23)
	assert.Nil(t, err)
	assert.Equal(t, "2016-04-30-23", name)
}

func Test2016043100IsNotValid(t *testing.T) {
	_, err := baseName(2016, 4, 31, 00)
	assert.NotNil(t, err)
}

func Test2016050100IsNotValid(t *testing.T) {
	_, err := baseName(2016, 5, 1, 00)
	assert.NotNil(t, err)
}

func Test2015010100IsValid(t *testing.T) {
	name, err := baseName(2015, 1, 1, 0)
	assert.Nil(t, err)
	assert.Equal(t, "2015-01-01-0", name)
}

func Test2014123123IsNotValid(t *testing.T) {
	_, err := baseName(2014, 12, 31, 23)
	assert.NotNil(t, err)
}
