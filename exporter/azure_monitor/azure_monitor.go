package azure_monitor

import (
	"errors"
	"bytes"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/trace"
	"fmt"
	"net/http"
 	"encoding/json"
)

// We can then use this for logs and trace exporter.
type Transporter struct {
	Envel Envelope
}

type Envelope struct { // TODO: Add more for next PR
	IKey string `json:"iKey"`
	Tags map[string]interface{} `json:"tags"`
	Name string `json:"name"`
} 


type AzureTraceExporter struct {
	projectID          string
	InstrumentationKey string
	options            common.Options
}

func NewAzureTraceExporter(o common.Options) (*AzureTraceExporter, error) {
	if o.InstrumentationKey == "" {
		return nil, errors.New("missing Instrumentation Key for Azure Exporter")
	}
	e := &AzureTraceExporter {
		//projectID:          "abcdefghijk",
		InstrumentationKey: o.InstrumentationKey,
		options:            o,
	}
	return e, nil
}

var _ trace.Exporter = (*AzureTraceExporter)(nil)

// Export SpanData to Azure Monitor
// The () before the function name means it is a function of AzureTraceExporter
func (e *AzureTraceExporter) ExportSpan(sd *trace.SpanData) {
	envelope := Envelope {
		IKey : e.options.InstrumentationKey,
		Tags : common.Azure_monitor_contect,
		Name : "Microsoft.ApplicationInsights.RemoteDependency",
	}

	transporter := Transporter{ 
		Envel: envelope,
	}
	transporter.Transmit(&e.options, envelope)
	// fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
	// 	sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}

func (e *Transporter) Transmit(o *common.Options, env Envelope) {
	fmt.Println("Begin Transmission") // For debugging
	//fmt.Println(env)
	fmt.Println(env.IKey)
	bytesRepresentation, err := json.Marshal(env)
	if err != nil {
		fmt.Println(err)
        fmt.Println("What happened?")
	}
	fmt.Println("Byte Representation")
	fmt.Println(string(bytesRepresentation))

	req, err := http.NewRequest("POST", o.EndPoint, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	req = req
	fmt.Println("REQUEST")
	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
        fmt.Println("What happened?")
	}
	resp = resp
	fmt.Println("RESPONSE")
	fmt.Println(resp)
}
