package main

import (
	"log"
	"net/http"

	"github.com/SazedWorldbringer/navigator-backend/routes"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/graph", http.HandlerFunc(routes.GraphHandler))

	corsMux := middlewareCors(mux)

	server := http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
