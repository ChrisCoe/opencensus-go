package main
// Package: Runs code for using Azure exporter

import (
	"context"
	
	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter := azure_monitor.NewAzureTraceExporter()
	exporter.InstrumentationKey = "11111111-1111-1111-1111-111111111111"

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
  
	_, span := trace.StartSpan(ctx, "/foo") // This calls the function ExportSpan written in azure_monitor.go 
	span.End()
}
