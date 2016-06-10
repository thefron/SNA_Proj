package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

func shouldNot(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var inputFileDir string
	var outputFileName string
	flag.StringVar(&inputFileDir, "input", "input", "input file directory")
	flag.StringVar(&outputFileName, "output", "output", "output file name")

	flag.Parse()
	fmt.Println("Start with input:", inputFileDir)
	fmt.Println("           output:", outputFileName)

	files, err := ioutil.ReadDir(inputFileDir)
	shouldNot(err)
	outputFile, err := os.Create(outputFileName)
	shouldNot(err)
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	lengthOfFiles := len(files)
	members := make(map[string](map[int]struct{}), lengthOfFiles)
	for _, file := range files {
		name, member := parse(inputFileDir, file.Name())
		members[name] = member
	}

	for i1, file1 := range files {
		for i2 := i1 + 1; i2 < lengthOfFiles; i2 += 1 {
			file2 := files[i2]

			langName1 := langName(file1.Name())
			member1, ok := members[langName1]
			if !ok {
				panic(errors.New(fmt.Sprintf("%s is not a valid name", langName1)))
			}
			langName2 := langName(file2.Name())
			member2, ok := members[langName2]
			if !ok {
				panic(errors.New(fmt.Sprintf("%s is not a valid name", langName2)))
			}

			writer.WriteString(fmt.Sprintf("%s %s", strings.Replace(langName1, " ", "-", -1), strings.Replace(langName2, " ", "-", -1)))

			common := common_distance(member1, member2)
			writer.WriteString(fmt.Sprintf(" common %d", common))

			jaccard := jaccard_distance(member1, member2)
			writer.WriteString(fmt.Sprintf(" jaccard %f", jaccard))
			writer.WriteByte('\n')
			writer.Flush()
		}
	}
}

func langName(fileName string) string {
	return fileName[0 : len(fileName)-len(".txt")]
}

func common_distance(lang1 map[int]struct{}, lang2 map[int]struct{}) int {
	c := 0
	for l1 := range lang1 {
		_, ok := lang2[l1]
		if ok {
			c += 1
		}
	}
	return c
}

func jaccard_distance(lang1 map[int]struct{}, lang2 map[int]struct{}) float64 {
	a := len(lang1)
	b := len(lang2)
	c := common_distance(lang1, lang2)
	return float64(a+b-c-c) / float64(a+b-c)
}

func parse(inputDir string, fileName string) (string, map[int]struct{}) {
	path := path.Join(inputDir, fileName)
	name := langName(fileName)

	input, err := os.Open(path)
	shouldNot(err)
	defer input.Close()

	reader := bufio.NewReader(input)
	scanner := bufio.NewScanner(reader)
	result := make(map[int]struct{}, 0)
	for scanner.Scan() {
		shouldNot(scanner.Err())
		id, err := strconv.Atoi(scanner.Text())
		shouldNot(err)
		result[id] = struct{}{}
	}

	return name, result
}
