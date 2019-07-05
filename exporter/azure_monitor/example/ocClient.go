package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

func main() {
	ctx := context.Background() // In other usages, the context would have been passed down after starting some traces.
	req, _ := http.NewRequest("GET", "https://opencensus.io/", nil)

	// It is imperative that req.WithContext is used to
	// propagate context and use it in the request.
	req = req.WithContext(ctx)

	client := &http.Client{Transport: &ochttp.Transport{}}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}

	// Consume the body and close it.
	io.Copy(ioutil.Discard, res.Body)
	_ = res.Body.Close()
}
