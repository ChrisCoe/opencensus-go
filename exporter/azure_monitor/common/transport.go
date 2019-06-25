package common
// Package: Structs commonly used for both trace and log exporters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// We can then use this for logs and trace exporter.
type Transporter struct {
	EnvelopeData Envelope
}

/*	Transmits envelope data to Azure Monitor.
	@param options holds specific attributes for exporter
	@param envelope Contains the data package to be transmitted
	@return The exporter created, and error if there is any
*/
func (e *Transporter) Transmit(options *Options, envelope *Envelope) {
	bytesRepresentation, err := json.Marshal(envelope)
	if err != nil {
		fmt.Println(err)
	}
	
	response, err := http.Post(
		options.EndPoint, 						//url
		"application/json",		 				//header
		bytes.NewBuffer(bytesRepresentation),	//data
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
