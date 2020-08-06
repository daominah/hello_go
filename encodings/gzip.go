package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mywrap/log"
)

type Response struct {
	RequestId string
	Rows      interface{}
}
type ResponseGzip struct {
	RequestId   string
	RowsGzipped []byte
}

func compressGzip(obj interface{}) ([]byte, error) {
	jsoned, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("json Marshal: %v", err)
	}
	var gzipped bytes.Buffer
	gzipWriter, err := gzip.NewWriterLevel(&gzipped, gzip.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("gzip NewWriterLevel: %v", err)
	}
	if _, err := gzipWriter.Write(jsoned); err != nil {
		return nil, fmt.Errorf("gzipWriter Write: %v", err)
	}
	if err := gzipWriter.Close(); err != nil {
		return nil, fmt.Errorf("gzipWriter Close: %v", err)
	}
	return gzipped.Bytes(), nil
}

func main() {
	var nRows = 100
	rows := make([]string, nRows)
	for i, _ := range rows {
		rows[i] = strings.Repeat("a", 100)
	}

	resp := Response{
		RequestId: "request0",
		Rows:      rows,
	}
	gzipped, err := compressGzip(resp.Rows)
	if err != nil {
		log.Fatalf("error compressGzip: %v", err)
	}
	respGzip := ResponseGzip{
		RequestId:   resp.RequestId,
		RowsGzipped: gzipped,
	}

	beauty1, _ := json.Marshal(respGzip)
	log.Printf("beauty1: len: %v, body: %s", len(beauty1), beauty1)
	beauty0, _ := json.Marshal(resp)
	log.Printf("beauty0: len: %v, body: %s", len(beauty0), beauty0[:100])
}
