package dijkstra

import (
	"math"
)

const infinity = math.MaxInt64

type Edge struct {
	To     int
	Weight int
}

type WeightedAdjacencyList [][]Edge

func hasUnseen(seen []bool, dists []int) bool {
	for i, v := range seen {
		if !v && dists[i] < infinity {
			return true
		}
	}
	return false
}

func getLowestUnseen(seen []bool, dists []int) int {
	idx := -1
	lowestDist := infinity

	for i, v := range seen {
		if v {
			continue
		}
		if lowestDist > dists[i] {
			lowestDist = dists[i]
			idx = i
		}
	}
	return idx
}

func DijkstraPath(source, sink int, graph WeightedAdjacencyList, intermediates []int) []int {
	seen := make([]bool, len(graph))

	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	dists := make([]int, len(graph))
	for i := range dists {
		dists[i] = infinity
	}
	dists[source] = 0

	for hasUnseen(seen, dists) {
		curr := getLowestUnseen(seen, dists)
		seen[curr] = true

		adjs := graph[curr]
		for _, edge := range adjs {
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	path := []int{}
	curr := sink
	for prev[curr] != -1 {
		path = append(path, curr)
		curr = prev[curr]
	}

	n := len(path)
	for i := 0; i < len(path)/2; i++ {
		path[i], path[n-1-i] = path[n-1-i], path[i]
	}

	return append([]int{source}, path...)
}
