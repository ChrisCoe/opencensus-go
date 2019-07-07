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
	// Line below only required if APPINSIGHTS_INSTRUMENTATIONKEY is not set
	//exporter.InstrumentationKey = "111a0d2f-ab53-4b62-a54f-4722f09fd136" 

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
  
	_, span := trace.StartSpan(ctx, "/foo") // This calls the function ExportSpan written in azure_monitor.go 
	span.End()
}
