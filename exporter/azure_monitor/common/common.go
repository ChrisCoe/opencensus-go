package common
// Package: Structs commonly used for both trace and log exporters

import (
	"context" 
	"fmt"
	"os"
)

var AzureMonitorContext = map[string]interface{} {
	"ai.cloud.role": "Go Application",
	"ai.cloud.roleInstance": getHostName(),
	"ai.device.id": getHostName(),
	"ai.device.type": "Other",
	"ai.internal.sdkVersion": "go:exp0.1",
}

func getHostName() (string) {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Problem with getting host name")
	}
	return hostName
}

type Options struct {
	ServiceName         string
	InstrumentationKey  string
	Context             context.Context
	EndPoint            string
	TimeOut             int
}

type Data struct {
	BaseData interface{}        `json:"baseData"`
	BaseType string             `json:"baseType"`
}

type Envelope struct {
	IKey string                 `json:"iKey"`
	Tags map[string]interface{} `json:"tags"`
	Name string                 `json:"name"`
	Time string                 `json:"time"`
	DataToSend Data             `json:"data"`
} 

type RemoteDependency struct {
	Name string         `json:"name"`
	Id string           `json:"id"`
	ResultCode string   `json:"resultCode"`
	Duration string     `json:"duration"`
	Success bool        `json:"success"`
	Ver int             `json:"ver"`
	Type string         `json:"type"`
}

type Request struct {
	Name string         `json:"name"`
	Id string           `json:"id"`
	Duration string     `json:"duration"`
	ResponseCode string `json:"responseCode"`
	Success bool        `json:"success"`
	Url string          `json:"url"`
}
