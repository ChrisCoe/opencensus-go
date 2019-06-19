package common

import (
	"context" 
)


type AzureMonitorContext struct {
    sdkVersion ai_internal_sdkVersion
}

type ai_internal_sdkVersion struct {
    goPlatform_version string
    opencensus_version string
    ext_version string
}

type BaseObject struct {
	ver int
	Name string
	time string
	sampleRate int
	success bool
	IKey string
	tags string
	data string
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

type Envelope struct {
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
