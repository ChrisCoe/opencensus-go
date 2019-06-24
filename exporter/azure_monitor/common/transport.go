package common

import (
	"bytes"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

// We can then use this for logs and trace exporter.
type Transporter struct {
	Envel Envelope
}

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
