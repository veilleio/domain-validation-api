// Domain Validation API – Go example
// Usage: VEILLE_API_KEY=your-key go run intelligence_domain.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client = &http.Client{Timeout: 120 * time.Second}

func call(method, target string, body []byte) []byte {
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req, _ := http.NewRequest(method, target, reader)
	req.Header.Set("x-api-key", os.Getenv("VEILLE_API_KEY"))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		fmt.Fprintf(os.Stderr, "HTTP %d: %s\n", resp.StatusCode, data)
		os.Exit(1)
	}
	return data
}

func main() {
	_ = url.Values{}
	_ = json.Valid
	params := url.Values{}
	params.Set("query", "tempmail.org")
	target := "https://api.veille.io/v1/intelligence/domain"
	if len(params) > 0 {
		target += "?" + params.Encode()
	}
	data := call("GET", target, nil)
	var pretty bytes.Buffer
	json.Indent(&pretty, data, "", "  ")
	fmt.Println(pretty.String())
}
