package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			// InsecureSkipVerify: true,
			InsecureSkipVerify: false,
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://golang.org/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Response: %+v\n", resp)
}

// https://stackoverflow.com/questions/12122159/how-to-do-a-https-request-with-bad-certificate
