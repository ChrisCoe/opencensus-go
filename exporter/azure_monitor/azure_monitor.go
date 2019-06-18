package azure_monitor

import (
	"fmt"
	"errors"
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/trace"
)


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
		projectID:          "abcdefghijk",
		InstrumentationKey: o.InstrumentationKey,
		options:            o,
	}
	return e, nil
}

var _ trace.Exporter = (*AzureTraceExporter)(nil)

//Export SpanData to Azure Monitor
// The () before the function name means it is a function of AzureTraceExporter
func (e *AzureTraceExporter) ExportSpan(sd *trace.SpanData) {
	baseObj := common.BaseObject {IKey : e.options.InstrumentationKey}
	envelope := common.Envelope {
		BaseObject: baseObj,
		//StartTime:  		sd.StartTime,
	}
	envelope.Name = "Microsoft.ApplicationInsights.RemoteDependency"

	transporter := common.Transporter{ 
		Envel: envelope,
	}
	transporter.Transmit(&e.options, &envelope)

	fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}
