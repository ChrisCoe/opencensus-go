package azure_monitor

import (
	"errors"
	"time"
	"fmt" // for debugging
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
	envelope := common.Envelope {
		IKey : e.options.InstrumentationKey,
		Tags : common.Azure_monitor_contect,
		Name : "Microsoft.ApplicationInsights.RemoteDependency",
		Time : getCurrentTime(),
	}

	if sd.ParentSpanID.String() == "" { // check this, it should enter if there is a parent span id
		fmt.Println("If statement is wrong")
		//envelope.Tags["ai.operation.parentId"] = 
	}
	if sd.SpanKind == trace.SpanKindServer {
		fmt.Println("ADD SERVER CASE")
		envelope.Name = "Microsoft.ApplicationInsights.Request"
	} else {
		fmt.Println("gucci")
		envelope.Name = "Microsoft.ApplicationInsights.RemoteDependency"
		currentData := common.RemoteDependency{
			Name : sd.Name,
			Id : "|" + sd.SpanContext.TraceID.String() + "." + sd.SpanID.String() + ".",
			ResultCode : "0", // not sure if needed,
			//Duration : sd.EndTime.Sub(sd.StartTime).String(), // TODO: might need to add more for formating
			Success : true,
		}
		if sd.SpanKind == trace.SpanKindClient {
			fmt.Println("add for this client case")
		} else {
			currentData.Type = "INPROC" //Check if this has an affect in golang
		}
		envelope.DataToSend = common.Data {
			BaseData : currentData,
			BaseType : "RemoteDependencyData",
		}
	}

	transporter := common.Transporter{ 
		Envel: envelope,
	}
	transporter.Transmit(&e.options, &envelope)
	// fmt.Printf("Name: %s\nTraceID: %x\nSpanID: %x\nParentSpanID: %x\nStartTime: %s\nEndTime: %s\nAnnotations: %+v\n\n",
	// 	sd.Name, sd.TraceID, sd.SpanID, sd.ParentSpanID, sd.StartTime, sd.EndTime, sd.Annotations)
}

/* Generates the current time stamp and properly formats
	@return time stamp
*/
func getCurrentTime() string {
	t := time.Now()
	formattedTime := t.Format("2006-01-02T15:04:05.000000Z")
	return formattedTime
}
