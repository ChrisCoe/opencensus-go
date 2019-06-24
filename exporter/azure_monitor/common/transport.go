package common
// Package: Structs commonly used for both trace and log exporters

import (
	"bytes"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
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
	fmt.Println("Begin Transmission") // For debugging
	bytesRepresentation, err := json.Marshal(envelope)
	if err != nil {
		fmt.Println("Error: json conversion for envelope\n")
	}
  
	reponse, err := http.Post(
		options.EndPoint, 						//url
		"application/json",		 				//header
		bytes.NewBuffer(bytesRepresentation),	//data
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
  
	log.Println(result)
	log.Println(result["data"])
	fmt.Println("End Transmission") // For debugging
}
