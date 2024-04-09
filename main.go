package main

import (
	"fmt"

	"github.com/SazedWorldbringer/navigator-backend/internal/dijkstra"
)

func main() {
	fmt.Println("Dijkstra's Algorithm")

	//      (1) --- (4) ---- (5)
	//    /  |       |       /|
	// (0)   | ------|------- |
	//    \  |/      |        |
	//      (2) --- (3) ---- (6)

	graph1 := dijkstra.WeightedAdjacencyList{
		{{To: 1, Weight: 3}, {To: 2, Weight: 1}},
		{{To: 0, Weight: 3}, {To: 2, Weight: 4}, {To: 4, Weight: 1}},
		{{To: 1, Weight: 4}, {To: 3, Weight: 7}, {To: 0, Weight: 1}},
		{{To: 2, Weight: 7}, {To: 4, Weight: 5}, {To: 6, Weight: 1}},
		{{To: 1, Weight: 1}, {To: 3, Weight: 5}, {To: 5, Weight: 2}},
		{{To: 6, Weight: 1}, {To: 4, Weight: 2}, {To: 2, Weight: 18}},
		{{To: 3, Weight: 1}, {To: 5, Weight: 1}},
	}

	source1 := 0
	sink1 := 6
	path1 := dijkstra.DijkstraPath(source1, sink1, graph1)

	fmt.Printf("Graph 1: %v\n", graph1)
	fmt.Printf("Source: %d\n", source1)
	fmt.Printf("Sink: %d\n", sink1)
	fmt.Printf("Path: %v\n", path1)

	//     >(1)<--->(4) ---->(5)
	//    /          |       /|
	// (0)     ------|------- |
	//    \   v      v        v
	//     >(2) --> (3) <----(6)

	graph2 := dijkstra.WeightedAdjacencyList{
		{{To: 1, Weight: 3}, {To: 2, Weight: 1}},
		{{To: 4, Weight: 1}},
		{{To: 3, Weight: 7}},
		{},
		{{To: 1, Weight: 1}, {To: 3, Weight: 5}, {To: 5, Weight: 2}},
		{{To: 6, Weight: 1}, {To: 2, Weight: 18}},
		{{To: 3, Weight: 1}},
	}

	source2 := 0
	sink2 := 6
	path2 := dijkstra.DijkstraPath(source1, sink1, graph1)

	fmt.Printf("Graph 1: %v\n", graph2)
	fmt.Printf("Source: %d\n", source2)
	fmt.Printf("Sink: %d\n", sink2)
	fmt.Printf("Path: %v\n", path2)
}
