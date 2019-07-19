package azure_monitor
// Package: extension for exporters to Azure Monitor.
// This includes examples on how to create azure exporters to send spans.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.opencensus.io/exporter/azure_monitor/utils"
	//"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/trace"
)

// Exporter is an implementation of trace.Exporter that uploads spans to Azure Monitor.
type Exporter struct {
	ServiceName         string
	InstrumentationKey  string
	EndPoint            string
	TimeOut             int
}

// Options are the options to be used when initializing an Azure Monitor exporter.
type Options struct {
	ServiceName         string
	InstrumentationKey  string	// required by user
	EndPoint            string
	TimeOut             int
}

// NewExporter returns an exporter that exports traces to Azure Monitor.
func NewExporter(o Options) (*Exporter, error) {
	if o.InstrumentationKey == "" {
		return nil, errors.New("missing Instrumentation Key for Azure Exporter")
	}
	if o.EndPoint == "" {
		o.EndPoint = "https://dc.services.visualstudio.com/v2/track"
	}
	if o.TimeOut == 0 {
		o.TimeOut = 10.0
	}
	e := &Exporter {
        ServiceName:         o.ServiceName,
        InstrumentationKey:  o.InstrumentationKey,
        Context:             o.Context,
        EndPoint:            o.EndPoint,
        TimeOut:             o.TimeOut,
	}
    return e, nil
}

// func NewAzureTraceExporter() (*AzureTraceExporter) {
// 	exporter := new(AzureTraceExporter)
// 	exporter.Options.EndPoint = "https://dc.services.visualstudio.com/v2/track"
// 	exporter.Options.TimeOut = 10.0
// 	return exporter
// }

var _ trace.Exporter = (*Exporter)(nil)

/*	Opencensus trace function required by interface. Called for every span/trace call.
	@param sd Span data retrieved by opencensus
*/
func (exporter *Exporter) ExportSpan(sd *trace.SpanData) {
	if exporter.InstrumentationKey == "" {
		log.Fatal(errors.New("missing Instrumentation Key for Azure Exporter"))
	}
	envelope := Envelope {
		IKey : exporter.InstrumentationKey,
		Tags : AzureMonitorContext,
		Time : utils.FormatTime(sd.StartTime),
	}
	
	envelope.Tags["ai.operation.id"] = sd.SpanContext.TraceID.String()
	if sd.ParentSpanID.String() != "0000000000000000" {
		envelope.Tags["ai.operation.parentId"] = "|" + sd.SpanContext.TraceID.String() + 
												 "." + sd.ParentSpanID.String()
	}
	if sd.SpanKind == trace.SpanKindServer {
		envelope.Name = "Microsoft.ApplicationInsights.Request"
		currentData := Request{
			Id : "|" + sd.SpanContext.TraceID.String() + "." + sd.SpanID.String() + ".",
			Duration : utils.TimeStampToDuration(sd.EndTime.Sub(sd.StartTime)),
			ResponseCode : "0",
			Success : true,
		}
		if _, isIncluded := sd.Attributes["http.method"]; isIncluded {
			currentData.Name = fmt.Sprintf("%s", sd.Attributes["http.method"])
		}
		if _, isIncluded := sd.Attributes["http.url"]; isIncluded {
			currentData.Name = fmt.Sprintf("%s %s", currentData.Name, sd.Attributes["http.url"])
			currentData.Url = fmt.Sprintf("%s", sd.Attributes["http.url"])
		}
		if _, isIncluded := sd.Attributes["http.status_code"]; isIncluded {
			currentData.ResponseCode = fmt.Sprintf("%d", sd.Attributes["http.status_code"])
		}
		envelope.DataToSend = Data {
			BaseData : currentData,
			BaseType : "RequestData",
		}

	} else {
		envelope.Name = "Microsoft.ApplicationInsights.RemoteDependency"
		currentData := RemoteDependency{
			Name : sd.Name,
			Id : "|" + sd.SpanContext.TraceID.String() + "." + sd.SpanID.String() + ".",
			ResultCode : "0", // TODO: Out of scope for now
			Duration : utils.TimeStampToDuration(sd.EndTime.Sub(sd.StartTime)),
			Success : true,
			Ver : 2,
		}
		if sd.SpanKind == trace.SpanKindClient {
			currentData.Type = "HTTP"
			if _, isIncluded := sd.Attributes["http.url"]; isIncluded {
				Url := fmt.Sprintf("%s", sd.Attributes["http.url"])
				currentData.Name = Url // TODO: parse URL before assignment
			}
			if _, isIncluded := sd.Attributes["http.status_code"]; isIncluded {
				currentData.ResultCode = fmt.Sprintf("%d", sd.Attributes["http.status_code"])
			}
		} else {
			currentData.Type = "INPROC" 
		}
		envelope.DataToSend = Data {
			BaseData : currentData,
			BaseType : "RemoteDependencyData",
		}
	}
	transporter := Transporter{ 
		EnvelopeData: envelope,
	}
	transporter.Transmit(exporter, &envelope)

	fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}
