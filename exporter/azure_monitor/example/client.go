package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: "11111111-1111-1111-1111-111111111111", // add your InstrumentationKey
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

	req, _ := http.NewRequest("GET", "http://localhost:8080/", nil)
	// It is imperative that req.WithContext is used to
	// propagate context and use it in the request.
	req = req.WithContext(ctx)
	client := &http.Client{Transport: &ochttp.Transport{}}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}
	// Consume the body and close it.
	io.Copy(ioutil.Discard, response.Body)
	_ = response.Body.Close()
}
