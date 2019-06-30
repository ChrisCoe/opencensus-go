package main
// Package: Structs commonly used for both trace and log exporters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	//bytesRepresentation, err := json.Marshal(envelope)

	hardCode := "{\"iKey\": \"111a0d2f-ab53-4b62-a54f-4722f09fd136\", \"tags\": {\"ai.cloud.role\": \"error.py\", \"ai.cloud.roleInstance\": \"MININT-1KVDB5T\", \"ai.device.id\": \"MININT-1KVDB5T\", \"ai.device.locale\": \"en_US\", \"ai.device.osVersion\": \"10.0.17763\", \"ai.device.type\": \"Other\", \"ai.internal.sdkVersion\": \"py3.7.3:oc0.6.0:ext0.2.0\", \"ai.operation.id\": \"00000000000000000000000000000000\", \"ai.operation.parentId\": \"|00000000000000000000000000000000.0000000000000000.\"}, \"time\": \"2019-06-30T08:57:52.315535Z\", \"name\": \"Microsoft.ApplicationInsights.Exception\", \"data\": {\"baseData\": {\"exceptions\": [{\"id\": 1, \"outerId\": 0, \"typeName\": \"ZeroDivisionError\", \"message\": \"Captured an exception.\\nTraceback (most recent call last):\\n  File \\\"error.py\\\", line 13, in main\\n    return 1" + `\/` + " 0  # generate a ZeroDivisionError\\nZeroDivisionError: division by zero\", \"hasFullStack\": true, \"parsedStack\": [{\"level\": 0, \"method\": \"main\", \"fileName\": \"error.py\", \"line\": 13}]}], \"severityLevel\": 3, \"properties\": {\"process\": \"MainProcess\", \"module\": \"error\", \"fileName\": \"error.py\", \"lineNumber\": 15, \"level\": \"ERROR\"}, \"ver\": 2}, \"baseType\": \"ExceptionData\"}}"
	response, err := http.Post(
		"https://dc.services.visualstudio.com/v2/track", 	//url
		"application/json",		 							//header
		//bytes.NewBuffer(bytesRepresentation),				//data
		bytes.NewBuffer([]byte(hardCode)),					//data
	)
	if err != nil {
		fmt.Println("Error: post error %d\n", err)
	}

	defer response.Body.Close() // prevent possible resource leak

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error: check decoder\n")
	}
}
