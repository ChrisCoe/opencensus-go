package main
// Package: Runs code for using Azure exporter

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"go.opencensus.io/stats/view"
)

func main() {
	// Firstly, we'll register ochttp Server views.
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatalf("Failed to register server views for HTTP metrics: %v", err)
	}
	// Enable observability to extract and examine stats.
	enableObservabilityAndExporters()
	// The handler containing your business logic to process requests.
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consume the request's body entirely.
		io.Copy(ioutil.Discard, r.Body)
		// Generate some payload of random length.
		res := strings.Repeat("aBa ", rand.Intn(20)+1)
		time.Sleep(time.Duration(rand.Intn(100)+1) * time.Millisecond)
		// Finally write the body to the response.
		w.Write([]byte("Hello, Chicken! " + res))
	})
	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}

	cst := httptest.NewServer(och)
	defer cst.Close()

	client := &http.Client{}
	for {
		body := strings.NewReader(strings.Repeat("aCa ", rand.Intn(10)+1))
		fmt.Println("urlBoy")
		fmt.Println(cst.URL)
		req, _ := http.NewRequest("POST", cst.URL, body)
		res, _ := client.Do(req)
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
		time.Sleep(979 * time.Millisecond)
	}
}

func enableObservabilityAndExporters() {
	exporter, err := azure_monitor.NewAzureTraceExporter("111a0d2f-ab53-4b62-a54f-4722f09fd136")
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
