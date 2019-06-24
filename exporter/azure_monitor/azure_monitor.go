package azure_monitor
// Package: extension for exporters to Azure Monitor.
// This includes examples on how to create azure exporters to send spans.

import (
	"errors"
	"time"
	"fmt" // for debugging
	"go.opencensus.io/exporter/azure_monitor/common"
	"go.opencensus.io/trace"
)

type AzureTraceExporter struct {
	ProjectID          string
	InstrumentationKey string
	Options            common.Options
}

/*	Creates an Azure Trace Exporter.
	@param options holds specific attributes for the new exporter
	@return The exporter created and error if there is any
*/
func NewAzureTraceExporter(Options common.Options) (*AzureTraceExporter, error) {
	if Options.InstrumentationKey == "" {
		return nil, errors.New("missing Instrumentation Key for Azure Exporter")
	}
	exporter := &AzureTraceExporter {
		ProjectID:          "abcdefghijk",
		InstrumentationKey: Options.InstrumentationKey,
		Options:            Options,
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
		Tags : common.Azure_monitor_contect,
		Name : "Microsoft.ApplicationInsights.RemoteDependency",
		Time : getCurrentTime(),
	}

	if sd.ParentSpanID.String() == "" { 
		// TODO: Add parent span details if any
		fmt.Println("ADD PARENT DETAILS")
		//envelope.Tags["ai.operation.parentId"] = 
	}
	if sd.SpanKind == trace.SpanKindServer {
		// TODO: Add for server case
		fmt.Println("ADD SERVER CASE")
		envelope.Name = "Microsoft.ApplicationInsights.Request"
	} else {
		envelope.Name = "Microsoft.ApplicationInsights.RemoteDependency"
		currentData := common.RemoteDependency{
			Name : sd.Name,
			Id : "|" + sd.SpanContext.TraceID.String() + "." + sd.SpanID.String() + ".",
			ResultCode : "0", // TODO: Out of scope for now
			Duration : timeStampToDuration(sd.EndTime.Sub(sd.StartTime)),
			Success : true,
			Ver : 2,
		}
		if sd.SpanKind == trace.SpanKindClient {
			// TODO: Add for client case
			fmt.Println("ADD CLIENT CASE")
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

/* Generates the current time stamp and properly formats to a string.
	@return time stamp
*/
func getCurrentTime() string {
	t := time.Now()
	formattedTime := t.Format("2006-01-02T15:04:05.000000Z")
	return formattedTime
}

/* Calculates number of days, hours, minutes, seconds, and milliseconds in a
	time duration. Then it properly formats into a string.
	@param t Time Duration
	@return formatted string 
*/
func timeStampToDuration(t time.Duration) (string) { 
	nanoSeconds := t.Nanoseconds()
	milliseconds, remainder :=  divMod(nanoSeconds, 1000000)
	seconds, remainder :=       divMod(remainder, 1000)
	minutes, remainder :=       divMod(remainder, 60)
	hours, remainder :=         divMod(remainder, 60)
	days, remainder :=          divMod(remainder, 24)
	
	formattedDays:=          fmt.Sprintf("%01d", days)
	formattedHours:=         fmt.Sprintf("%02d", hours)
	formattedMinutes :=      fmt.Sprintf("%02d", minutes)
	formattedSeconds :=      fmt.Sprintf("%02d", seconds)
	formattedMilliseconds := fmt.Sprintf("%03d", milliseconds)

	return formattedDays + "." + formattedHours + ":" + formattedMinutes + ":" + formattedSeconds + "."+ formattedMilliseconds
}

/* Performs division and returns both quotient and remainder. */
func divMod(numerator, denominator int64) (quotient, remainder int64) {
    quotient = numerator / denominator // integer division, decimals are truncated
    remainder = numerator % denominator
    return
}
