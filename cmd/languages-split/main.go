package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func shouldNot(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var inputFileName string
	var outputFilePrefix string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&outputFilePrefix, "output", "output", "output prefix")
	flag.Parse()
	fmt.Println("Start with input:", inputFileName)
	fmt.Println("           output:", outputFilePrefix)

	inputFile, err := os.Open(inputFileName)
	shouldNot(err)
	defer inputFile.Close()

	reader := bufio.NewReader(inputFile)
	_, err = reader.ReadString(' ')
	shouldNot(err)
	for {
		for {
			b, err := reader.ReadByte()
			if err == io.EOF {
				return
			}
			shouldNot(err)
			if b == '<' {
				break
			}
		}
		language := ""
		for {
			b, err := reader.ReadByte()
			shouldNot(err)
			if b == '>' {
				break
			}
			language = fmt.Sprintf("%s%c", language, b)
		}
		_, err = reader.ReadByte()
		shouldNot(err)

		outputFileName := fmt.Sprintf("%s%s.txt", outputFilePrefix, language)
		outputFile, err := os.Create(outputFileName)
		shouldNot(err)
		defer outputFile.Close()
		writer := bufio.NewWriter(outputFile)
		user := ""
		for {
			b, err := reader.ReadByte()
			shouldNot(err)
			if b == ' ' || b == '\n' || b == '\r' {
				_, err = writer.WriteString(user)
				user = ""
				shouldNot(err)
				_, err = writer.WriteString("\n")
				shouldNot(err)
				writer.Flush()
				if b == '\n' || b == '\r' {
					break
				}
				continue
			}
			user = fmt.Sprintf("%s%c", user, b)
		}
		writer.Flush()
	}
}
