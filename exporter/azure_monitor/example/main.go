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
		InstrumentationKey: "111111111111111111", // add your InstrumentationKey
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
