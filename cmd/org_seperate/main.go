package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func shouldNot(e error) {
	if e != nil {
		panic(e)
	}
}

type Node int
type nodes map[Node]int

func (n nodes) add(i Node) {
	value, ok := n[i]
	if ok {
		n[i] = value + 1
	} else {
		n[i] = 1
	}
}

type edge struct {
	id int
	n1 Node
	n2 Node
}
type edges map[edge]struct{}

func min(n1 Node, n2 Node) Node {
	if n1 > n2 {
		return n2
	}
	return n1
}
func max(n1 Node, n2 Node) Node {
	if n1 < n2 {
		return n2
	}
	return n1
}
func (e edges) add(n1 Node, n2 Node, id int) {
	minN := min(n1, n2)
	maxN := max(n1, n2)
	e[edge{id, minN, maxN}] = struct{}{}
}

func main() {
	var inputFileName string
	var userNumberoforgsFileName string
	var orgNumberofusersFileName string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&userNumberoforgsFileName, "user-numberoforgs", "user-numberoforgs.txt", "user-numberoforgs file name")
	flag.StringVar(&orgNumberofusersFileName, "org-numberofusers", "org-numberofusers.txt", "org-numberofusers file name")
	flag.Parse()
	fmt.Println("Start with input:", inputFileName)
	fmt.Println("           user-numberoforgs:", userNumberoforgsFileName)
	fmt.Println("           org-numberofusers:", orgNumberofusersFileName)

	inputFile, err := os.Open(inputFileName)
	shouldNot(err)
	defer inputFile.Close()
	userNumberoforgsFile, err := os.Create(userNumberoforgsFileName)
	shouldNot(err)
	defer userNumberoforgsFile.Close()
	orgNumberofusersFile, err := os.Create(orgNumberofusersFileName)
	shouldNot(err)
	defer orgNumberofusersFile.Close()

	nodes := make(nodes, 100)
	edges := make(edges, 100)

	scanner := bufio.NewScanner(inputFile)
	orgNumberofusersWriter := bufio.NewWriter(orgNumberofusersFile)
	defer orgNumberofusersWriter.Flush()
	for {
		if !scanner.Scan() {
			shouldNot(scanner.Err())
			break
		}
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		id, err := strconv.Atoi(tokens[0])
		shouldNot(err)
		lineToWrite := fmt.Sprintf("%d %d\n", id, len(tokens)-2)
		_, err = orgNumberofusersWriter.WriteString(lineToWrite)
		shouldNot(err)
		orgNumberofusersWriter.Flush()
		members := []int{}
		for i := 2; i < len(tokens); i += 1 {
			member, err := strconv.Atoi(tokens[i])
			shouldNot(err)
			members = append(members, member)
			nodes.add(Node(member))
		}

		for _, target := range members {
			for _, neighbor := range members {
				if target == neighbor {
					continue
				}
				edges.add(Node(target), Node(neighbor), id)
			}
		}
	}

	userNumberoforgsWriter := bufio.NewWriter(userNumberoforgsFile)
	defer userNumberoforgsWriter.Flush()
	for node, count := range nodes {
		line := fmt.Sprintf("%d %d\n", node, count)
		_, err := userNumberoforgsWriter.WriteString(line)
		shouldNot(err)
		userNumberoforgsWriter.Flush()
	}
}
