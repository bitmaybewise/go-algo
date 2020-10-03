package algo

import (
	"bytes"
	"fmt"
	"math"
)

type undirectedGraph struct {
	vertices []byte
	adjList  map[byte][]byte
}

type weightedGraph struct {
	vertices []byte
	adjList  map[byte]map[byte]int
}

func NewGraph() *undirectedGraph {
	return new(undirectedGraph)
}

func NewWeightedGraph() *weightedGraph {
	return new(weightedGraph)
}

func (g *undirectedGraph) AddVertex(vertex byte) {
	g.vertices = append(g.vertices, vertex)
	if g.adjList == nil {
		g.adjList = make(map[byte][]byte)
	}
	g.adjList[vertex] = make([]byte, 0)
}

func (g *weightedGraph) AddVertex(vertex byte) {
	g.vertices = append(g.vertices, vertex)
	if g.adjList == nil {
		g.adjList = make(map[byte]map[byte]int)
	}
}

func (g *undirectedGraph) AddEdge(v byte, w byte) {
	g.adjList[w] = append(g.adjList[w], v)
	g.adjList[v] = append(g.adjList[v], w)
}

func (g *weightedGraph) AddEdge(vertex byte, edge byte, weight int) {
	if g.adjList[vertex] == nil {
		g.adjList[vertex] = make(map[byte]int)
	}
	g.adjList[vertex][edge] = weight

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

func (g *weightedGraph) String() string {
	var b bytes.Buffer

	for _, vertex := range g.vertices {
		for adjVertex, weight := range g.adjList[vertex] {
			b.WriteString(string(adjVertex) + " -> " + fmt.Sprint(weight) + " -> ")
		}

		b.WriteString("\n")
	}

	return b.String()
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

func (g *weightedGraph) findLowestCostVertex(costs map[byte]int, processed map[byte]bool) (lowestCostVertex byte) {
	lowestCost := math.MaxInt64

	for vertex, weight := range costs {
		if weight < lowestCost && !processed[vertex] {
			lowestCost = weight
			lowestCostVertex = vertex
		}
	}

	return lowestCostVertex
}

// Dijkstra returns a path with the minimum number of edges using Dijkstra's search.
func (g *weightedGraph) Dijkstra(a byte, b byte) ([]byte, int) {
	costs, parents := make(map[byte]int), make(map[byte]byte)

	for neighbor, weight := range g.adjList[a] {
		costs[neighbor] = weight
		parents[neighbor] = a
	}
	if _, ok := costs[b]; !ok {
		costs[b] = math.MaxInt64
	}

	processed := make(map[byte]bool)
	node := g.findLowestCostVertex(costs, processed)
	for node != 0 {
		cost := costs[node]
		for neighbor, weight := range g.adjList[node] {
			newCost := cost + weight
			if costs[neighbor] > newCost {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}
		processed[node] = true
		node = g.findLowestCostVertex(costs, processed)
	}

	path, to := make([]byte, 0), b
	for to != 0 {
		path = append(path, to)
		to = parents[to]
	}
	reversedPath := make([]byte, len(path))
	for i, v := range path {
		reversedPath[len(path)-i-1] = v
	}

	return reversedPath, costs[b]
}
