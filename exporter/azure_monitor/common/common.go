package common

import (
	"context" 
	"runtime"
	//"strings"
)

// Some of these structs will be used in future PRs

var Azure_monitor_contect = map[string]interface{} {
	"ai.cloud.role": "Golang Application",
	"ai.internal.sdkVersion":  runtime.Version() + ":oc" + Opencensus_version + ":ext" + Ext_version,
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
	Name string
	IKey string
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
	baseDate string
	baseType string
}

type Envelope struct { // TODO: Add more for next PR
	BaseObject
}

type RemoteDependency struct {
	BaseObject
	id string
	duration string
	responseCode string
	url string
	properties string
	measurements string
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
