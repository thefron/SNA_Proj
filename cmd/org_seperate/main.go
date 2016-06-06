package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
func (e edges) from(n Node) []edge {
	edges := []edge{}
	for edge, _ := range e {
		if edge.n1 != n {
			continue
		}
		edges = append(edges, edge)
	}
	return edges
}

type visited map[Node]struct{}

func (v visited) visit(n Node) {
	v[n] = struct{}{}
}
func (v visited) isVisited(n Node) bool {
	_, ok := v[n]
	return ok
}

type Edges []edge

func (e Edges) Len() int {
	return len(e)
}

func (e Edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e Edges) Less(i, j int) bool {
	ei := e[i]
	ej := e[j]
	if ei.id < ej.id {
		return true
	}
	if ei.id > ej.id {
		return false
	}
	if ei.n1 < ej.n1 {
		return true
	}
	if ei.n1 > ej.n1 {
		return false
	}
	return ei.n2 < ej.n2
}

func main() {
	var inputFileName string
	var userNumberoforgsFileName string
	var orgNumberofusersFileName string
	var communityFileNamePrefix string
	flag.StringVar(&inputFileName, "input", "input.txt", "input file name")
	flag.StringVar(&userNumberoforgsFileName, "user-numberoforgs", "user-numberoforgs.txt", "user-numberoforgs file name")
	flag.StringVar(&orgNumberofusersFileName, "org-numberofusers", "org-numberofusers.txt", "org-numberofusers file name")
	flag.StringVar(&communityFileNamePrefix, "community-prefix", "community", "community file name prefix")
	flag.Parse()
	fmt.Println("Start with input:", inputFileName)
	fmt.Println("           user-numberoforgs:", userNumberoforgsFileName)
	fmt.Println("           org-numberofusers:", orgNumberofusersFileName)
	fmt.Println("           community-prefix:", communityFileNamePrefix)

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

	nodeArrays := []int{}
	for node, _ := range nodes {
		nodeArrays = append(nodeArrays, int(node))
	}
	sort.IntSlice(nodeArrays).Sort()

	visited := make(visited, len(nodes))
	filePostfix := 0
	for i, node := range nodeArrays {
		if visited.isVisited(Node(node)) {
			continue
		}
		fmt.Println(i, "of", len(nodeArrays), len(nodes))
		community := make([]edge, 0)

		candidates := edges.from(Node(node))
		handled := make(map[edge]struct{}, 0)
		for candidateIndex := 0; candidateIndex < len(candidates); candidateIndex += 1 {
			candidate := candidates[candidateIndex]
			community = append(community, candidate)
			visited.visit(candidate.n2)
			es := edges.from(candidate.n2)
			for _, e := range es {
				_, ok := handled[candidate]
				if !ok {
					handled[candidate] = struct{}{}
					candidates = append(candidates, e)
				}
			}
		}

		if len(community) != 0 {
			communityFileName := fmt.Sprintf("%s%d.txt", communityFileNamePrefix, filePostfix)
			filePostfix += 1
			communityFile, err := os.Create(communityFileName)
			shouldNot(err)

			communityWriter := bufio.NewWriter(communityFile)
			sort.Sort(Edges(community))
			for _, edge := range community {
				line := fmt.Sprintf("%d %d %d\n", edge.id, edge.n1, edge.n2)
				_, err := communityWriter.WriteString(line)
				shouldNot(err)
				communityWriter.Flush()
			}
			communityFile.Close()
		}
	}
}
