package main

import (
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	exporter, err := azure_monitor.NewAzureTraceExporter("11111111-1111-1111-1111-111111111111")
	if err != nil {
		log.Fatal(err)
	}
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}

	// Now use the instrumented handler
	if err := http.ListenAndServe(":8080", och); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
