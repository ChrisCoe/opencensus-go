package azure_monitor

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
	options            common.Options
}

func NewAzureTraceExporter(o common.Options) (*AzureTraceExporter, error) {
	if o.InstrumentationKey == "" {
		return nil, errors.New("missing Instrumentation Key for Azure Exporter")
	}
	e := &AzureTraceExporter {
		ProjectID:          "abcdefghijk",
		InstrumentationKey: o.InstrumentationKey,
		options:            o,
	}
	return e, nil
}

var _ trace.Exporter = (*AzureTraceExporter)(nil)

// Export SpanData to Azure Monitor
// The () before the function name means it is a function of AzureTraceExporter
func (e *AzureTraceExporter) ExportSpan(sd *trace.SpanData) {
	envelope := common.Envelope {
		IKey : e.options.InstrumentationKey,
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
		Envel: envelope,
	}
	transporter = transporter
	transporter.Transmit(&e.options, &envelope)

	fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
		sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}

/* Generates the current time stamp and properly formats
	@return time stamp
*/
func getCurrentTime() string {
	t := time.Now()
	formattedTime := t.Format("2006-01-02T15:04:05.000000Z")
	return formattedTime
}

/* Calcuates number of days, hours, minutes, seconds, and miliseconds of a
	time duration. Then it properly formats into a string.
	@param t Time Duration
	@return formatted string 
*/
func timeStampToDuration(t time.Duration) (string) { 
	nanoSeconds := t.Nanoseconds()
	miliseconds, remainder := 	divmod(nanoSeconds, 1000000)
	seconds, remainder := 		divmod(remainder, 1000)
	minutes, remainder := 		divmod(remainder, 60)
	hours, remainder := 		divmod(remainder, 60)
	days, remainder := 			divmod(remainder, 24)

	formattedDays:=  		 fmt.Sprintf("%01d", days)
	formattedHours:=  		 fmt.Sprintf("%02d", hours)
	formattedMinutes :=  	 fmt.Sprintf("%02d", minutes)
	formattedSeconds :=  	 fmt.Sprintf("%02d", seconds)
	formattedMiliseconds :=  fmt.Sprintf("%03d", miliseconds)

	return formattedDays + "." + formattedHours + ":" + formattedMinutes + ":" + formattedSeconds + "."+ formattedMiliseconds
}

/* Performs division and returns both quotient and remainder */
func divmod(numerator, denominator int64) (quotient, remainder int64) {
    quotient = numerator / denominator // integer division, decimals are truncated
    remainder = numerator % denominator
    return
}
