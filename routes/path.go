package routes

import (
	"encoding/json"
	"net/http"

	"github.com/SazedWorldbringer/navigator-backend/internal/dijkstra"
)

type pathReq struct {
	Source int                            `json:"source"`
	Sink   int                            `json:"sink"`
	Graph  dijkstra.WeightedAdjacencyList `json:"graph"`
}

func ShortestPathHandler(w http.ResponseWriter, r *http.Request) {
	var req pathReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path := dijkstra.DijkstraPath(req.Source, req.Sink, req.Graph)
	jsonResponse, err := json.Marshal(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
