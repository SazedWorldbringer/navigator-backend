package dijkstra

import (
	"fmt"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		source, sink int
		graph        WeightedAdjacencyList
		expected     []int
	}

	graph1 := WeightedAdjacencyList{
		{{1, 3}, {2, 1}},
		{{0, 3}, {2, 4}, {4, 1}},
		{{1, 4}, {3, 7}, {0, 1}},
		{{2, 7}, {4, 5}, {6, 1}},
		{{1, 1}, {3, 5}, {5, 2}},
		{{6, 1}, {4, 2}, {2, 18}},
		{{3, 1}, {5, 1}},
	}

	graph2 := WeightedAdjacencyList{
		{{1, 3}, {2, 1}},
		{{4, 1}},
		{{3, 7}},
		{},
		{{1, 1}, {3, 5}, {5, 2}},
		{{6, 1}, {2, 18}},
		{{3, 1}},
	}

	tests := []testCase{
		{
			source:   0,
			sink:     6,
			graph:    graph1,
			expected: []int{0, 1, 4, 5, 6},
		},
		{
			source:   0,
			sink:     6,
			graph:    graph2,
			expected: []int{0, 1, 4, 5, 6},
		},
	}

	for _, test := range tests {
		if actual := DijkstraPath(test.source, test.sink, test.graph); !slices.Equal(actual, test.expected) {
			t.Errorf(`Test Failed:
source: %d
sink: %d
graph: %v
=>
actual path: %v
expected path: %v
`, test.source, test.sink, test.graph, actual, test.expected)
		} else {
			fmt.Printf(`Test Passed:
source: %d
sink: %d
graph: %v
=>
actual path: %v
expected path: %v
`, test.source, test.sink, test.graph, actual, test.expected)
		}
	}
}
