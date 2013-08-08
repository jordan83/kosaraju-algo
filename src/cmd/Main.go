package main

import (
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
	"strconv"
	"../pkg/graph"
	"../pkg/sccs"
	"sort"
)

func readEdge(graph graphs.Graph, line string) (tail, head *graphs.Vertex) {
	components := strings.Split(line, " ")
	
	tailLabel, _ := strconv.Atoi(components[0])
	headLabel, _ := strconv.Atoi(components[1])
	
	tail = graph.GetVertex(tailLabel -1)
	head = graph.GetVertex(headLabel -1)
	return
}

func addEdges(graph graphs.Graph, edges []string) {
	for _,edge := range edges {
		tail, head := readEdge(graph, edge)
		graph.AddNeighbour(tail, head) 
	}
}

func initGraph(edges []string) graphs.Graph {
	maxStrLabel := strings.Split(edges[len(edges) -1], " ")[0]	
	maxLabel, _ := strconv.Atoi(maxStrLabel)
	
	graph := graphs.NewGraph(maxLabel)
	addEdges(graph, edges)
	
	return graph
}

func readInputLines() (lines []string, err error) {
	fileName := flag.String("fileName", "SCC.txt", "File name of the graph to parse")
	flag.Parse()
	
	content, err := ioutil.ReadFile(*fileName)
	if err != nil {
		lines = []string{}
		return
	}
	
	contentStr := string(content)
	lines = strings.Split(strings.TrimSpace(contentStr), "\n")
	return
}

func printLargestSccs(l scc.SccList) {
	sort.Sort(l)
	count := 0
	for _, v := range l.Items {
		fmt.Printf("Leader: %v, Size: %v\n", v.Leader, v.Size)
		count = count + 1
		if count >= 5 {
			break
		}
	}
}

func main() {
	lines,_ := readInputLines()
	graph := initGraph(lines)
	
	l := scc.FindSccs(graph)
	printLargestSccs(l)
}