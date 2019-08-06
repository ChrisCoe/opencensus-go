package azure_monitor
// Package: extension for exporters to Azure Monitor.
// This includes examples on how to create azure exporters to send spans.

import (
	"errors"
	"fmt"

	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/exporter/azure_monitor/utils"
	"go.opencensus.io/trace"
)

type AzureTraceExporter struct {
	InstrumentationKey string
	Options            common.Options
}

/*	Creates an Azure Trace Exporter.
	@param options holds specific attributes for the new exporter
	@return The exporter created and error if there is any
*/
func NewAzureTraceExporter(IKey string) (*AzureTraceExporter, error) {
	if IKey == "" {
		return nil, errors.New("missing Instrumentation Key for Azure Exporter")
	}
	currentOptions := common.Options {
		InstrumentationKey: IKey,
		EndPoint:           "https://dc.services.visualstudio.com/v2/track",
		TimeOut: 			10.0,
	}
	exporter := &AzureTraceExporter {
		InstrumentationKey: currentOptions.InstrumentationKey,
		Options:            currentOptions,
	}
	
	return exporter, nil
}

var _ trace.Exporter = (*AzureTraceExporter)(nil)

/*	Opencensus trace function required by interface. Called for every span/trace call.
	@param sd Span data retrieved by opencensus
*/
func (exporter *AzureTraceExporter) ExportSpan(sd *trace.SpanData) {
	envelope := common.Envelope {
		IKey : exporter.Options.InstrumentationKey,
		Tags : common.AzureMonitorContext,
		Time : utils.FormatTime(sd.StartTime),
	}

	envelope.Tags["ai.operation.id"] = sd.SpanContext.TraceID.String()
	if sd.ParentSpanID.String() != "0000000000000000" {
		envelope.Tags["ai.operation.parentId"] = "|" + sd.SpanContext.TraceID.String() + "." + sd.ParentSpanID.String() + "."
	}
	if sd.SpanKind == trace.SpanKindServer {
		envelope.Name = "Microsoft.ApplicationInsights.Request"
		currentData := common.Request{
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
		envelope.DataToSend = common.Data {
			BaseData : currentData,
			BaseType : "RequestData",
		}

	} else {
		envelope.Name = "Microsoft.ApplicationInsights.RemoteDependency"
		currentData := common.RemoteDependency{
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
		envelope.DataToSend = common.Data {
			BaseData : currentData,
			BaseType : "RemoteDependencyData",
		}
	}
	transporter := common.Transporter{ 
		EnvelopeData: envelope,
	}
	transporter.Transmit(&exporter.Options, &envelope)

	fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}
