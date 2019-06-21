package common

// import (
// 	"bytes"
// 	//"log"
// 	"net/http"
// 	"encoding/json"
// 	"fmt"
// 	//"io/ioutil"
// )

// // We can then use this for logs and trace exporter.
// type Transporter struct {
// 	Envel Envelope
// }

// func (e *Transporter) Transmit(o *common.Options, env Envelope) {
// 	fmt.Println("Begin Transmission") // For debugging
// 	//fmt.Println(env)
// 	fmt.Println(env.IKey)
// 	bytesRepresentation, err := json.Marshal(env)
// 	if err != nil {
// 		fmt.Println(err)
//         fmt.Println("What happened?")
// 	}
// 	fmt.Println("Byte Representation")
// 	fmt.Println(string(bytesRepresentation))

// 	req, err := http.NewRequest("POST", o.EndPoint, bytes.NewBuffer(bytesRepresentation))
// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Set("Accept", "application/json")

// 	req = req
// 	fmt.Println("REQUEST")
// 	fmt.Println(req)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
//         fmt.Println("What happened?")
// 	}
// 	resp = resp
// 	fmt.Println("RESPONSE")
// 	fmt.Println(resp)
// }
