package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: "11111111-1111-1111-1111-111111111111", // add your InstrumentationKey
		EndPoint: 			"https://dc.services.visualstudio.com/v2/track",
		TimeOut: 			10.0,
	})
	if err != nil {
		log.Fatal(err)
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
  
	ctx, span := trace.StartSpan(ctx, "/parent") // This calls the function ExportSpan written in azure_monitor.go 
	boo(ctx)
	span.End()
}

func boo(ctx context.Context) {
	ctx, span := trace.StartSpan(ctx, "/child")
	defer span.End()

	response, err := http.Get("http://localhost:8080/")
	if err != nil {
			log.Fatal(err)
	}
}