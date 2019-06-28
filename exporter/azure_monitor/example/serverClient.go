package main

import (

	"io"
	"io/ioutil"
	"math/rand"
	"net/http/httptest"
	"strings"
	"time"

	"log"
	"net/http"

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
		res := strings.Repeat("a", rand.Intn(99971)+1)

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
		body := strings.NewReader(strings.Repeat("a", rand.Intn(777)+1))
		req, _ := http.NewRequest("POST", cst.URL, body)
		res, _ := client.Do(req)
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
		time.Sleep(979 * time.Millisecond)
	}

	// // Now use the instrumnted handler
	// if err := http.ListenAndServe(":8080", och); err != nil {
	// 	log.Fatalf("Failed to run the server: %v", err)
	// }
}

func enableObservabilityAndExporters() {
	exporter, err := azure_monitor.NewAzureTraceExporter("111a0d2f-ab53-4b62-a54f-4722f09fd136")
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}
