package common
// Package: Structs commonly used for both trace and log exporters

import (
	"context" 
)

type AzureMonitorContext struct {
    SdkVersion ai_internal_sdkVersion
}

type ai_internal_sdkVersion struct {
    GoPlatform_version string
    Opencensus_version string
    Ext_version string
}

type BaseObject struct { // Used to avoid repeat attributes
	Version int
	Name string
	Time string
	SampleRate int
	Success bool
	IKey string
	Tags string
	Data string
}

type Options struct {
	ServiceName        	string
	InstrumentationKey 	string
	Context            	context.Context
	EndPoint			string
	TimeOut				int
}

type Data struct {
	BaseDate string
	BaseType string
}

type Envelope struct { //TODO: Add more attributes
	BaseObject
}

type RemoteDependency struct {
	BaseObject
	Id string
	Duration string
	ResponseCode string
	Url string
	Properties string
	Measurements string
}

type Request struct {
	BaseObject
	Id string
	Duration string
	ResponseCode string
	Url string
	Properties string
	Measurements string
}
