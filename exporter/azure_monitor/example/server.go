
package main



import (

	"context"

	"log"

	"net/http"



	"go.opencensus.io/exporter/azure_monitor"

	"go.opencensus.io/exporter/azure_monitor/common"

	"go.opencensus.io/plugin/ochttp"

	"go.opencensus.io/trace"

)



func main() {

	originalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, Chicken!"))



		ctx := context.Background()



		exporter, err := azure_monitor.NewAzureTraceExporter(
			InstrumentationKey: "11111111-1111-1111-1111-111111111111" // add your InstrumentationKey
		)

		if err != nil {

			log.Fatal(err)

		}



		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

		trace.RegisterExporter(exporter)

	

		_, span := trace.StartSpan(ctx, "/serverSide") // This calls the function ExportSpan written in azure_monitor.go 

		span.End() //TODO: Investigate why this span is not considered trace.SpanKindServer
	})

	och := &ochttp.Handler{
		Handler: originalHandler, // The handler you'd have used originally
	}

	// Now use the instrumnted handler
	if err := http.ListenAndServe(":8080", och); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}