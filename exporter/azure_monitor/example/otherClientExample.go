package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/exporter/azure_monitor/utils"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background() // In other usages, the context would have been passed down after starting some traces.
	utils.enableObservabilityAndExporter()
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
	fmt.Println(res)
}
