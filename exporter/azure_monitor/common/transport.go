package common

import (
	"bytes"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	//"time"
	//"io/ioutil"
)

// We can then use this for logs and trace exporter.
type Transporter struct {
	Envel Envelope
	
}

func (e *Transporter) Transmit(options *Options, env *Envelope) {
	fmt.Println("Begin Transmission\n") // For debugging
	//fmt.Println(env)
	//fmt.Println(time.Now().UTC())
	hardCode := "[{\"iKey\": \"d07ba4f7-7546-47b4-b3e0-7fa203f17f6a\", \"tags\": {\"ai.cloud.role\": \"simple.py\", \"ai.cloud.roleInstance\": \"MININT-1KVDB5T\", \"ai.device.id\": \"MININT-1KVDB5T\", \"ai.device.locale\": \"en_US\", \"ai.device.osVersion\": \"10.0.17763\", \"ai.device.type\": \"Other\", \"ai.internal.sdkVersion\": \"py3.7.3:oc0.6.0:ext0.2.0\", \"ai.operation.id\": \"60cf1cb9518d0ba3a71856ffa81ac05a\"}, \"time\": \"2019-06-22T20:50:59.812277Z\", \"name\": \"Microsoft.ApplicationInsights.RemoteDependency\", \"data\": {\"baseData\": {\"name\": \"potato_a\", \"id\": \"|60cf1cb9518d0ba3a71856ffa81ac05a.9c0907432b1294a1.\", \"resultCode\": \"0\", \"duration\": \"0.00:00:00.000\", \"success\": true, \"ver\": 2, \"type\": \"INPROC\"}, \"baseType\": \"RemoteDependencyData\"}}]" // noice
	fmt.Println("hey listen!")
	fmt.Println([]byte(hardCode))
	bytesRepresentation, err := json.Marshal(env)
	if err != nil {
		fmt.Println(err)
        fmt.Println("What happened?")
	}
	fmt.Println("Byte Representation")
	fmt.Println(string(bytesRepresentation))

	
	reponse, err := http.Post(
		options.EndPoint, 						//url
		"application/json",		 				//header
		bytes.NewBuffer([]byte(hardCode)),	//data
	)
	if err != nil {
		fmt.Println("Error: post error %d\n", err)
	}

	defer reponse.Body.Close() // prevent possible resource leak

	var result map[string]interface{}
	err = json.NewDecoder(reponse.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error: check decoder\n")
	}

	fmt.Println("\nRESULT")
	log.Println(result)
	log.Println(result["data"])

	// req, err := http.NewRequest("POST", o.EndPoint, bytes.NewBuffer(bytesRepresentation))
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Accept", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
    //     fmt.Println("What happened?")
	// }

	fmt.Println("\nRESPONSE")
	fmt.Println(reponse)

	fmt.Println("End Transmission") // For debugging
}
