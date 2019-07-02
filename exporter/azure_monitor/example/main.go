package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"log"
	
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: "111a0d2f-ab53-4b62-a54f-4722f09fd136", // add your InstrumentationKey
		EndPoint: 			"https://dc.services.visualstudio.com/v2/track",
		TimeOut: 			10.0,
	})
	if err != nil {
		log.Fatal(err)
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
  
	_, span := trace.StartSpan(ctx, "/foo") // This calls the function ExportSpan written in azure_monitor.go 
	span.End()
}
