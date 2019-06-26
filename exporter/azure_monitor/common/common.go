package common
// Package: Structs commonly used for both trace and log exporters

import (
	"context" 
	"runtime"
	"os"
	"fmt"
)

var Azure_monitor_contect = map[string]interface{} {
	"ai.cloud.role": "Go Application",
	"ai.cloud.roleInstance": getHostName(),
	"ai.device.id": getHostName(),
	"ai.device.type": "Other",
	"ai.internal.sdkVersion":  runtime.Version() + ":oc" + Opencensus_version + ":ext" + Ext_version,
}

func getHostName() (string) {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Problem with getting host name")
	}
	return hostName
}

type Options struct {
	ServiceName        	string
	InstrumentationKey 	string
	Context            	context.Context
	EndPoint			string
	TimeOut				int
}

type Data struct {
	BaseData interface{} 			`json:"baseData"`
	// BaseRequest Request 			`json:"baseData"`
	BaseType string 			`json:"baseType"`
}

type Envelope struct { // TODO: Add more for next PR
	IKey string 				`json:"iKey"`
	Tags map[string]interface{} `json:"tags"`
	Name string 				`json:"name"`
	Time string 				`json:"time"`
	DataToSend Data 			`json:"data"`
} 

type RemoteDependency struct {
	Name string 		`json:"name"`
	Id string 			`json:"id"`
	ResultCode string 	`json:"resultCode"`
	Duration string 	`json:"duration"`
	Success bool 		`json:"success"`
	Ver int				`json:"ver"`
	Type string			`json:"type"`
}

type Request struct {
	Name string 		`json:"name"`
	Id string 			`json:"id"`
	Duration string 	`json:"duration"`
	ResponseCode string `json:"responsecode"`
	Success bool 		`json:"success"`
	Url interface{} 	`json:"url"`
}
