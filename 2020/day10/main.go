package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var (
	pathCount int
)

// https://github.com/sohamkamani/golang-graph-traversal

// Vertex represents a Vertex used in DirectedGraph
type Vertex struct {
	Key      int
	Vertices map[int]*Vertex
}

// NewVertex returns a new vertex
func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

// DirectedGraph represents a Directed Graph
type DirectedGraph struct {
	Vertices map[int]*Vertex
	directed bool
}

// NewDirectedGraph returns a DirectedGraph
func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		Vertices: map[int]*Vertex{},
	}
}

// AddVertex adds a Vertex to the Graph
func (g *DirectedGraph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

// AddEdge adds an edge to the Graph
func (g *DirectedGraph) AddEdge(k1, k2 int) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		panic("not all vertices exist")
	}

	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2
}

// Sort returns a topological sorting of the graph
func (g *DirectedGraph) Sort() []*Vertex {
	orphans := []*Vertex{}
	sorted := []*Vertex{}
	indegree := map[*Vertex]int{}

	// 1. compute indegree
	for _, v := range g.Vertices {
		if _, ok := indegree[v]; !ok {
			indegree[v] = 0
		}
		for _, v2 := range v.Vertices {
			if _, ok := indegree[v2]; !ok {
				indegree[v2] = 1
			} else {
				indegree[v2]++
			}
		}
	}

	// 2. collect indegree = 0
	for v, id := range indegree {
		if id == 0 {
			orphans = append(orphans, v)
		}
	}

	// 3. sort
	for len(orphans) > 0 {
		current := orphans[len(orphans)-1]
		orphans = orphans[:len(orphans)-1]
		sorted = append(sorted, current)

		for _, v := range current.Vertices {
			indegree[v]--
			if indegree[v] == 0 {
				orphans = append(orphans, v)
			}
		}
	}

	return sorted
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// to string
	lines := strings.Split(string(data), "\n")

	adapters := []int{0}
	position := 0
	onediff := 0
	threediff := 0
	graph := NewDirectedGraph()
	graph.AddVertex(0)

	for _, line := range lines {
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		adapters = append(adapters, value)
		graph.AddVertex(value)
	}

	sort.Ints(adapters)

	// solution #1
	for {
		adapter := adapters[position]
		next := adapters[position+1]
		switch next - adapter {
		case 1:
			onediff++
			break
		case 2:
			break
		case 3:
			threediff++
			break
		default:
			panic("no adapter found")
		}

		position++

		if position > len(adapters)-2 {
			break
		}
	}

	// device's builtin adapter always 3 higher than the highest
	threediff++

	fmt.Printf("Solution #1: %d (1-diff: %d, 3-diff: %d)\n", onediff*threediff, onediff, threediff)

	// Solution #2
	// create a graph
	position = 0
	for {
		i := position + 1
		if i > len(adapters)-1 {
			break
		}
		for {
			if adapters[i]-adapters[position] <= 3 {
				graph.AddEdge(adapters[position], adapters[i])
			} else {
				break
			}
			i++
			if i > len(adapters)-1 {
				break
			}
		}
		position++
	}

	// we use a topological sort and then go backwards to get the number of paths
	sorted := graph.Sort()
	paths := map[int]int{}
	paths[adapters[len(adapters)-1]] = 1

	for key := len(sorted) - 1; key >= 0; key-- {
		if len(sorted[key].Vertices) == 0 {
			paths[adapters[key]] = 1
		}
		for _, edge := range sorted[key].Vertices {
			// for each outgoing edge, add to the path count
			paths[adapters[key]] += paths[edge.Key]

		}
	}

	fmt.Printf("Solution #2: %v\n", paths[0])
}
