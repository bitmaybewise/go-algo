package algo

import (
	"fmt"
	"testing"
)

func TestGraphShortestPath(t *testing.T) {
	graph := NewGraph()
	vertices := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	graph.AddEdge('A', 'B')
	graph.AddEdge('A', 'C')
	graph.AddEdge('A', 'D')
	graph.AddEdge('C', 'D')
	graph.AddEdge('C', 'G')
	graph.AddEdge('D', 'G')
	graph.AddEdge('D', 'H')
	graph.AddEdge('B', 'E')
	graph.AddEdge('B', 'F')
	graph.AddEdge('E', 'I')

	tests := []struct {
		from         byte
		to           byte
		expectedPath []byte
	}{
		{from: 'A', to: 'B', expectedPath: []byte{'A', 'B'}},
		{from: 'A', to: 'G', expectedPath: []byte{'A', 'C', 'G'}},
		{from: 'A', to: 'I', expectedPath: []byte{'A', 'B', 'E', 'I'}},
		{from: 'H', to: 'F', expectedPath: []byte{'H', 'D', 'A', 'B', 'F'}},
	}
	for _, s := range tests {
		t.Run(fmt.Sprintf("From %c to %c", s.from, s.to), func(t *testing.T) {
			trail := graph.ShortestPath(s.from, s.to)
			if len(trail) != len(s.expectedPath) {
				t.Errorf("Path should contain %d vertices", len(s.expectedPath))
			}
			for i, v := range trail {
				if v != s.expectedPath[i] {
					t.Errorf("Vertex at position %d should be %c but returned %c", i, s.expectedPath[i], v)
				}
			}
		})
	}
}

func TestGraphDijkstra(t *testing.T) {
	graph := NewWeightedGraph()
	vertices := []byte{'A', 'B', 'C', 'D'}
	for _, v := range vertices {
		graph.AddVertex(v)
	}
	graph.AddEdge('A', 'B', 2)
	graph.AddEdge('A', 'C', 6)
	graph.AddEdge('B', 'C', 3)
	graph.AddEdge('C', 'D', 1)
	graph.AddEdge('B', 'D', 5)

	tests := []struct {
		from         byte
		to           byte
		expectedPath []byte
		distance     int
	}{
		{from: 'A', to: 'B', expectedPath: []byte{'A', 'B'}, distance: 2},
		{from: 'A', to: 'C', expectedPath: []byte{'A', 'B', 'C'}, distance: 5},
		{from: 'B', to: 'C', expectedPath: []byte{'B', 'C'}, distance: 3},
		{from: 'A', to: 'D', expectedPath: []byte{'A', 'B', 'C', 'D'}, distance: 6},
	}
	for _, s := range tests {
		t.Run(fmt.Sprintf("From %c to %c", s.from, s.to), func(t *testing.T) {
			trail, distance := graph.Dijkstra(s.from, s.to)
			if len(trail) != len(s.expectedPath) {
				t.Errorf("Path should contain %d vertices", len(s.expectedPath))
			}
			if distance != s.distance {
				t.Errorf("Distance should be %d but returned %d", s.distance, distance)
			}
			for i, v := range trail {
				if v != s.expectedPath[i] {
					t.Errorf("Vertex at position %d should be %c but returned %c", i, s.expectedPath[i], v)
				}
			}
		})
	}
}
