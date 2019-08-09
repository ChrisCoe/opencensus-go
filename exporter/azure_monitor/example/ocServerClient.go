package main

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"go.opencensus.io/exporter/azure_monitor"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func main() {

	// Enable observability to extract and examine traces.
	enableObservabilityAndExporters()

	// The handler containing your business logic to process requests.
	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Consume the request's body entirely.
		io.Copy(ioutil.Discard, r.Body)
		
		// Generate some payload of random length.
		res := strings.Repeat("a", rand.Intn(99971)+1)
		
		// Sleep for a random time to simulate a real server's operation.
		time.Sleep(time.Duration(rand.Intn(977)+1) * time.Millisecond)

		// Finally write the body to the response.
		w.Write([]byte("Hello, World! " + res))
	})
	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}
	cst := httptest.NewServer(och)
	defer cst.Close()

	client := &http.Client{}
	for i := 0; i < 3; i++ {
		body := strings.NewReader(strings.Repeat("a", rand.Intn(777)+1))
		req, _ := http.NewRequest("POST", cst.URL, body)
		res, _ := client.Do(req)
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
		time.Sleep(979 * time.Millisecond)
	}
}

func enableObservabilityAndExporters() {
	exporter := azure_monitor.NewAzureTraceExporter()
	exporter.Options.InstrumentationKey = "111a0d2f-ab53-4b62-a54f-4722f09fd136"
	
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})  // why not always sample? I would miss some errors...
	trace.RegisterExporter(exporter)
}