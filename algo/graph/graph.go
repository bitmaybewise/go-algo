package graph

import (
	"bytes"
)

type undirectedGraph struct {
	vertices []byte
	adjList  map[byte][]byte
}

func (g *undirectedGraph) AddVertex(vertex byte) {
	g.vertices = append(g.vertices, vertex)
	if g.adjList == nil {
		g.adjList = make(map[byte][]byte)
	}
	g.adjList[vertex] = make([]byte, 0)
}

func (g *undirectedGraph) AddEdge(v byte, w byte) {
	g.adjList[w] = append(g.adjList[w], v)
	g.adjList[v] = append(g.adjList[v], w)
}

func (g *undirectedGraph) String() string {
	var b bytes.Buffer

	for _, v := range g.vertices {
		b.WriteString(string(v) + " -> ")

		for _, n := range g.adjList[v] {
			b.WriteString(string(n) + " ")
		}

		b.WriteString("\n")
	}

	return b.String()
}

func NewGraph() *undirectedGraph {
	return new(undirectedGraph)
}

// ShortestPath returns a path with the minimum number of edges using breadth-first search.
func (g *undirectedGraph) ShortestPath(a byte, b byte) []byte {
	visited := make(map[byte]bool)
	predecessors := make(map[byte]byte)
	for _, v := range g.vertices {
		visited[v] = false
		predecessors[v] = 0
	}

	queue := make([]byte, 1)
	queue[0] = a

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		visited[u] = true

		for _, w := range g.adjList[u] {
			if !visited[w] {
				visited[w] = true
				predecessors[w] = u
				queue = append(queue, w)
			}
		}
	}

	path, to := make([]byte, 0), b
	for to != 0 {
		path = append(path, to)
		to = predecessors[to]
	}
	reversedPath := make([]byte, len(path))
	for i, v := range path {
		reversedPath[len(path)-i-1] = v
	}

	return reversedPath
}
