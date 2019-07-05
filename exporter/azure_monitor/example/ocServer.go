package main

import (
	"log"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

func main() {
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}

	// Now use the instrumnted handler
	if err := http.ListenAndServe(":8080", och); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
