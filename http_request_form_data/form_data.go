package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
)

const (
	SPLITTER = "\n____________________________________________________________\n"
)

func SendFile() {
	// create request
	requestBody := &bytes.Buffer{}
	mapFieldToFilePath := map[string]string{
		"file1": "file1.jpg",
		"file2": "file2.txt",
	}
	writer := multipart.NewWriter(requestBody)
	for key, filePath := range mapFieldToFilePath {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fw, err := writer.CreateFormFile(key, file.Name())
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			panic(err)
		}

	}
	writer.Close()

	var requestUrl string
	//requestUrl = "http://localhost:8080"
	requestUrl = "http://localhost:8080/cms/v1/tickets/export?Page=1&PageSize=5"
	request, err := http.NewRequest("GET", requestUrl, requestBody)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4NjM4NTQsImlhdCI6MTU0NzE4Mzg1NCwidXNlcl9uYW1lIjoidi50dW5nZHQxMUB2aW5pZC5uZXQiLCJ1c2VyX2lkIjoyLCJwZXJtaXNzaW9ucyI6eyJBcHBWZXJzaW9uIjpbMSwyLDMsNCw1XSwiQXR0ZW5kZWUiOlsxLDgsOV0sIkNhY2hlIjpbMSwyLDNdLCJDbXNMb2ciOlsxXSwiRXZlbnQiOlsxLDQsNSw2LDldLCJFeGNlcHRpb25Mb2ciOlsxXSwiRXhwb3J0TG9nIjpbMV0sIkltcG9ydExvZyI6WzFdLCJNZXJjaGFudCI6WzEsMiwzLDQsNSw2XSwiTW9iaWxlTG9nIjpbMV0sIk5vdGlmaWNhdGlvbiI6WzEsMiwzLDQsNV0sIk9yZGVyIjpbMSw4LDldLCJPdGhlclByb21vdGlvbiI6WzEsMiwzLDQsNSw2LDddLCJQZXJtaXNzaW9uIjpbMSwyXSwiUHJvbW90aW9uIjpbMSwyLDMsNCw1LDYsN10sIlByb3ZpZGVyIjpbMSw1LDldLCJSZXNlcnZlIjpbMSw0LDUsNiw3LDgsMTEsMTJdLCJSb2xlIjpbMSwyLDMsNCw1XSwiU3RhZmYiOlsxLDIsMyw0LDUsNiw3LDhdLCJTdG9yZSI6WzEsMiwzLDQsNSw2XSwiVGFnIjpbMSwyLDMsNCw1LDZdLCJUaWNrZXQiOlsxLDUsOF0sIlRpY2tldFR5cGUiOlsxLDQsNSw2LDldLCJVc2VyIjpbMSwyLDNdLCJWb3VjaGVyIjpbMSwyLDMsNCw1LDYsNyw4XSwiVm91Y2hlclRpY2tldCI6WzEsMiwzLDQsNSw2XX19.CyIPzNHojub-puJgwiw_Mp5t3oHdAd0xbY8PLWA9FOq7vWnsSuIa0wNMSW9TuDj3gUvlwF_dzyoRU3Ku9Ikv3E4U7pbjam3UfN4TU17G1qcM0WpocdL47B8KenBu4IK5D_DwcT8lIBd8s-Dz2snqn_kvOIhX-Bw5a73JV-EFagTMgYojOpOcoV7d3qwcq_pED16X2r8UhBO5_xiG6CbX6MmywmN0ScOS0sLkvV2loroN-O04srPzRiXqMAi4w8HFgJ6V3Cvt1Xtf_M68houHmIjOhT2SrX_gy_Lu19l4xT1jxUwpn97V2ffcA0TSLYvyZg1O3Hn_slWc_mQ2edQnPUirOcavJODUJtryJV6rf-5Ow-Q4otS7KfgGtgyj4iPHdwmquIzAyMgtvRngxFujlFSOcpQ5qTYXIB3pICrgjbDuuqdUcHb0H5J8VMKqTBsl6BFogBMYx1mITXzU8OhJ_gAry35GST5PofqxHSW9RrYTy2UFOe7_6C4Fcjq7XFAIzy90Q60Jz7KryR-jIUZwPuEFCUwbMczRRE2spEkK9a7ZDWu7ugQYfVtFdT3T-lvgrE5m0vKX2hciojOo5Clo2FQMiDvHPrnwrNUbbf6R7V_37-JiUfXKzkrZlvLNVEqNZI_oPAcVTRTV8lHwS2sOnsWGB9eFimVvFQjw0J4vVEs")

	// print request
	dumpedRequest, err := httputil.DumpRequestOut(request, true)
	fmt.Printf("Err: %v. DumpedRequest:\n%v%v",
		err, string(dumpedRequest), SPLITTER)

	// send request
	response, err := (&http.Client{}).Do(request)
	if err != nil {
		panic(err)
	}

	// print response
	dumpedResponse, err := httputil.DumpResponse(response, true)
	fmt.Printf("Err: %v. DumpedResponse:\n%v%v",
		err, string(dumpedResponse), SPLITTER)
}
