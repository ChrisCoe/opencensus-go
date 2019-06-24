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
func (e *Transporter) Transmit(options *Options, env *Envelope) {
	fmt.Println("Begin Transmission\n") // For debugging
	bytesRepresentation, err := json.Marshal(env)
	if err != nil {
		fmt.Println(err)
        fmt.Println("What happened?")
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
  
	log.Println(result)
	log.Println(result["data"])

	fmt.Println("End Transmission") // For debugging
}
