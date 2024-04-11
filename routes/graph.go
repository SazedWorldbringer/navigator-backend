package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/SazedWorldbringer/navigator-backend/internal/dijkstra"
)

type reqBody struct {
	Vertices int `json:"vertices"`
	Edges    int `json:"edges"`
}

func GenerateRandomGraph(vertices, edges int) dijkstra.WeightedAdjacencyList {
	graph := make(dijkstra.WeightedAdjacencyList, vertices)
	for i := range graph {
		graph[i] = make([]dijkstra.Edge, 0)
	}

	for i := 0; i < edges; i++ {
		from := rand.Intn(vertices)
		to := rand.Intn(vertices)
		weight := rand.Intn(10) + 1

		if from != to && !contains(graph[from], to) {
			graph[from] = append(graph[from], dijkstra.Edge{To: to, Weight: weight})
		}
	}
	return graph
}

func contains(edges []dijkstra.Edge, to int) bool {
	for _, edge := range edges {
		if edge.To == to {
			return true
		}
	}
	return false
}

func GraphHandler(w http.ResponseWriter, r *http.Request) {
	var req reqBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	graph := GenerateRandomGraph(req.Vertices, req.Edges)
	jsonResponse, err := json.Marshal(graph)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
