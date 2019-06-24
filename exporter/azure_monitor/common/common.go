package common

import (
	"context" 
	"runtime"
	"os"
	"fmt"
	//"strings"
)

// Some of these structs will be used in future PRs

var Azure_monitor_contect = map[string]interface{} {
	"ai.cloud.role": "main.go",
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

type AzureMonitorContext struct {
    sdkVersion ai_internal_sdkVersion
}

type ai_internal_sdkVersion struct {
    goPlatform_version string
    opencensus_version string
    ext_version string
}

type BaseObject struct { // Used to avoid repeat attributes
	//ver int
	Name string `json:"name"`
	iKey string
	//time string
	//sampleRate int
	//success bool 
	Tags map[string]interface{}
	//data string
}

type Options struct {
	ServiceName        	string
	InstrumentationKey 	string
	Context            	context.Context
	EndPoint			string
	TimeOut				int
}

type Data struct {
	BaseData RemoteDependency `json:"baseData"`
	BaseType string `json:"baseType"`
}

type Envelope struct { // TODO: Add more for next PR
	IKey string `json:"iKey"`
	Tags map[string]interface{} `json:"tags"`
	Name string `json:"name"`
	Time string `json:"time"`
	DataToSend Data `json:"data"`
} 
type RemoteDependency struct {
	Name string 		`json:"name"`
	Id string 			`json:"id"`
	ResultCode string 	`json:"resultCode"`
	Duration string 	`json:"duration"`
	Success bool 		`json:"success"`
	Type string			`json:"type"`
	// responseCode string
	// url string
	// properties string
	// measurements string
}

type Request struct {
	BaseObject
	id string
	duration string
	responseCode string
	url string
	properties string
	measurements string
}
