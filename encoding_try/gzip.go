package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mywrap/log"
	"github.com/mywrap/textproc"
)

// obj can be a string
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
	s0 := textproc.GenRandomWord(1000000, 1000000)
	s1 := strings.Repeat("0123456789abcdefghik1011121314lmnopqrstu", 25000)
	s2Bld := strings.Builder{}
	for i := 0; true; i++ {
		s2Bld.WriteString(fmt.Sprintf("%v", i))
		if s2Bld.Len() >= 1000000 {
			break
		}
	}
	s2 := s2Bld.String()

	for _, origin := range []string{s0, s1, s2} {
		gzipped, _ := compressGzip(origin)
		log.Printf("fullString length: origin: %v, gzipped: %v, ratio: %.3f",
			len(origin), len(gzipped), float64(len(origin))/float64(len(gzipped)))
	}
}
