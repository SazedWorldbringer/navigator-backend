package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/SazedWorldbringer/navigator-backend/internal/dijkstra"
)

type graphReq struct {
	Vertices int `json:"vertices"`
}

func generateRandomGraph(vertices int, prob float64) dijkstra.WeightedAdjacencyList {
	graph := make(dijkstra.WeightedAdjacencyList, vertices)
	for i := range graph {
		graph[i] = make([]dijkstra.Edge, 0)
	}

	for from := 0; from < vertices; from++ {
		for to := from + 1; to < vertices; to++ {
			if from != to && rand.Float64() < prob {
				weight := rand.Intn(10) + 1
				graph[from] = append(graph[from], dijkstra.Edge{To: to, Weight: weight})
			}
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
	var req graphReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	graph := generateRandomGraph(req.Vertices, 0.2)
	jsonResponse, err := json.Marshal(graph)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
