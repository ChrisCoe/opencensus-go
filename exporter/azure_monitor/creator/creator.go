package creator
// Package: Function used to create exporters

import (
	"log"
	
	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/trace"
)

/* Create and set exporter for Azure Monitor */
func EnableObservabilityAndExporter() {
	exporter, err := azure_monitor.NewAzureTraceExporter("111a0d2f-ab53-4b62-a54f-4722f09fd136")
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
