## Introduction

Azure Monitor helps developers and customers to monitor their application.
This project aims to support all OpenCensus SDKs piping data to Azure Monitor through the OneAgent.

This package lets Golang developers send traces to Azure Monitor from OpenCEnsus Go SDK. 

**Project is still in very early development. This is not meant to be used for any production code, yet.**

## Examples

### How to create exporter and start a span
```go
ctx := context.Background()
exporter, err := azure_monitor.NewAzureTraceExporter(common.Options{
  // add your InstrumentationKey
  InstrumentationKey: "11111111-1111-1111-1111-111111111111",
  // end point for Azure Monitor to digest
  EndPoint: 			    "https://dc.services.visualstudio.com/v2/track",
  TimeOut: 			      10.0,
})
if err != nil {
  log.Fatal(err)
}

// For production, you might want to change sampling rate
trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
// Required to send span with Azure Monitor exporter
trace.RegisterExporter(exporter)

// StartSpan calls function ExportSpan which all exporters implement
_, span := trace.StartSpan(ctx, "/foo") 
span.End() // All spans need to end
```

### Server and client span uses

Send a span for client actions by using the OpenCensus http wrapper transport. Once you create your
handler just add it to the opencensus http wrapper as seen below. See https://godoc.org/go.opencensus.io/plugin/ochttp#Transport

```go
client := &http.Client{
    Transport: &ochttp.Transport{}
}
```

Send a span for server actions by using the OpenCensus http wrapper handler. Once you create your
handler just add it to the opencensus http wrapper as seen below. See https://godoc.org/go.opencensus.io/plugin/ochttp#Handler

```go
och := &ochttp.Handler{
    Handler: originalHandler, // The handler you'd have used originally
}
```


