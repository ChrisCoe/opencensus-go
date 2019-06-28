package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azure_monitor.NewAzureTraceExporter("111a0d2f-ab53-4b62-a54f-4722f09fd136")
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
	fmt.Println("start BOO")
	ctx, span := trace.StartSpan(ctx, "/child")
	defer span.End()

	// response, err := http.Get("http://localhost:8080/")
	// if err != nil {
			
	// 		log.Fatal(err)
	// }
	// fmt.Println(response)
	// fmt.Println("end BOO")
}
