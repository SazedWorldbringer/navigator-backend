package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SazedWorldbringer/navigator-backend/internal/dijkstra"
)

type pathReq struct {
	Source int    `json:"source"`
	Sink   int    `json:"sink"`
	Graph  string `json:"graph"`
}

func ShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	var req pathReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	graph := dijkstra.WeightedAdjacencyList{}
	err = json.Unmarshal([]byte(req.Graph), &graph)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	path := dijkstra.DijkstraPath(req.Source, req.Sink, graph, []int{})
	jsonResponse, err := json.Marshal(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
