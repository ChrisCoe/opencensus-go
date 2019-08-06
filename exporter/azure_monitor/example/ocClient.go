package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {
	/* In other usages, the context would have been passed down after starting some traces. */
	ctx := context.Background() 
	
	exporter := azure_monitor.NewAzureTraceExporter()
	exporter.Options.InstrumentationKey = "111a0d2f-ab53-4b62-a54f-4722f09fd136"
	
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})  // why not always sample? I would miss some errors...
	trace.RegisterExporter(exporter)

	/* This calls the function ExportSpan written in azure_monitor.go  */
	ctx, span := trace.StartSpan(ctx, "/parentGood2") 
	foo(ctx)
	span.End()
	log.Print("Program Terminated")
}

/* Function must take a context.Context as a parameter to create a child span
for the trace, which is a tree of spans.
*/
func foo(ctx context.Context) {
	ctx, span := trace.StartSpan(ctx, "/childGood2") // should be a child span
	defer span.End()

	req, _ := http.NewRequest("GET", "https://en.wikipedia.org/wiki/Chicken", nil)

	// It is imperative that req.WithContext is used to
	// propagate context and use it in the request.
	req = req.WithContext(ctx)

	client := &http.Client{Transport: &ochttp.Transport{}}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}

	// Consume the body and close it.
	io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()
}
