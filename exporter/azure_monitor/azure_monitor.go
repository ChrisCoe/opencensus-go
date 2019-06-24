package azure_monitor
// Package: extension for exporters to Azure Monitor.
// This includes examples on how to create azure exporters to send spans.

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

/*	Creates an Azure Trace Exporter.
	@param options holds specific attributes for the new exporter
	@return The exporter created and error if there is any
*/
func NewAzureTraceExporter(options common.Options) (*AzureTraceExporter, error) {
	if options.InstrumentationKey == "" {
		return nil, errors.New("missing Instrumentation Key for Azure Exporter")
	}
	exporter := &AzureTraceExporter {
		projectID:          "abcdefghijk",
		InstrumentationKey: options.InstrumentationKey,
		options:            options,
	}
	return exporter, nil
}

var _ trace.Exporter = (*AzureTraceExporter)(nil)

/*	Opencensus trace function required by interface. Called for every span/trace call.
	@param sd Span data retrieved by opencensus
*/
func (exporter *AzureTraceExporter) ExportSpan(sd *trace.SpanData) {
	baseObj := common.BaseObject {IKey : exporter.options.InstrumentationKey}
	envelope := common.Envelope {
		BaseObject: baseObj,
	}
	envelope.Name = "Microsoft.ApplicationInsights.RemoteDependency"

	transporter := common.Transporter{ 
		EnvelopeData: envelope,
	}
	transporter.Transmit(&exporter.options, &envelope)

	fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}
