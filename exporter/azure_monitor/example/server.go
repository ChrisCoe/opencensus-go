package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
			InstrumentationKey: "111a0d2f-ab53-4b62-a54f-4722f09fd136", // add your InstrumentationKey
			EndPoint: 			"https://dc.services.visualstudio.com/v2/track",
			TimeOut: 			10.0,
		})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte("Hello, Chicken!"))
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
		trace.RegisterExporter(exporter)

		_, span := trace.StartSpan(ctx, "/serverSide") // This calls the function ExportSpan written in azure_monitor.go 

		span.End()
	})

	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}

	// Now use the instrumnted handler
	if err := http.ListenAndServe(":8080", och); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
