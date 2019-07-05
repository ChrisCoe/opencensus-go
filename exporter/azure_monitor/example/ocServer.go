package main

import (
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
			InstrumentationKey: "11111111-1111-1111-1111-111111111111", // add your InstrumentationKey
		})
		if err != nil {
			log.Fatal(err)
		}
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
		trace.RegisterExporter(exporter)
	})
	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}

	// Now use the instrumented handler
	if err := http.ListenAndServe(":8080", och); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
