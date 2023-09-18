package main

import (
	"github.com/extism/go-pdk"
)

func main() {
	// create an HTTP Request (without relying on WASI), set headers as needed
	req := pdk.NewHTTPRequest("GET", "https://jsonplaceholder.typicode.com/todos/1")
	req.SetHeader("some-name", "some-value")
	req.SetHeader("another", "again")
	// send the request, get response back (can check status on response via res.Status())
	res := req.Send()

	// zero-copy output to host
	pdk.OutputMemory(res.Memory())
}
