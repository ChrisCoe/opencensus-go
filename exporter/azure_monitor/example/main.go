package main

import (
	"context"
	"log"
	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: "d07ba4f7-7546-47b4-b3e0-7fa203f17f6a", // add your InstrumentationKey
		EndPoint: 			"https://dc.services.visualstudio.com/v2/track",
		TimeOut: 			3.0,
	})
	if err != nil {
		log.Fatal(err)
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)

	_, span := trace.StartSpan(ctx, "/foo") // This calls the function ExportSpan written in azure_monitor.go 
	span.End()
}

// maybe we need ssl?