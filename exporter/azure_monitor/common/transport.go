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

func (e *Transporter) Transmit(o *Options, env *Envelope) {
	fmt.Println("Begin Transmission") // For debugging

	// env2 := map[string]interface{}{
	// 	"InstrumentationKey": env.BaseObject.IKey,
	// }

	fmt.Println(env.BaseObject)
	//fmt.Println(env2)
	bytesRepresentation, err := json.Marshal(env.BaseObject)
	// if err != nil {
	// 	fmt.Println("Error: json conversion for envelope\n")
	// }
	//fmt.Println(bytesRepresentation)

	// header := map[string]interface{}{
	// 	"Accept": "application/json",
	// }
	//var jsonStr = []byte(`{"instrumentation key":"d07ba4f7-7546-47b4-b3e0-7fa203f17f6a"}`)
	
	url := o.EndPoint
	response, err := http.Post(
		url, 							//url
		"application/json; charset=utf-8",		 				//header
		bytes.NewBuffer(bytesRepresentation),	//data
	)



	// reponse, err := http.Get(
	// 	url			,				//url
	// )
	// if err != nil {
	// 	fmt.Println("Error: post error %d\n", err)
	// }




	defer response.Body.Close() // prevent possible resource leak

	var result map[string]interface{}
	
	fmt.Println(response)
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error: check decoder\n")
	}

	log.Println(result)
	log.Println(result["data"])
	fmt.Println("End Transmission") // For debugging
}
