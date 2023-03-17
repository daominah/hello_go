package main

import (
	"bytes"
	"compress/gzip"
	cryptorand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gzWriter, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("gzip.NewWriterLevel: %v", err)
	}
	if _, err := gzWriter.Write(data); err != nil {
		return nil, fmt.Errorf("gzWriter.Write: %v", err)
	}
	if err := gzWriter.Close(); err != nil {
		return nil, fmt.Errorf("gzWriter.Close: %v", err)
	}
	return buf.Bytes(), nil
}

func uncompress(data []byte) ([]byte, error) {
	gzReader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("gzip.NewReader: %v", err)
	}
	ret, err := io.ReadAll(gzReader)
	if err != nil {
		return nil, fmt.Errorf("gzReader ReadAll: %v", err)
	}
	if err := gzReader.Close(); err != nil {
		return ret, fmt.Errorf("gzReader Close: %v", err)
	}
	return ret, nil
}

func main() {
	randomBytes := make([]byte, 1000000)
	_, _ = cryptorand.Read(randomBytes)
	for _, c := range []struct {
		caseName   string
		toCompress []byte
	}{
		{caseName: "repeated string", toCompress: []byte(strings.Repeat("0123456789abcdefghij9753124680klmnopqrst", 25000))},
		{caseName: "random string  ", toCompress: []byte(genRandomString(1000000))},
		{caseName: "random bytes   ", toCompress: randomBytes},
	} {
		gzipped, err := compress(c.toCompress)
		if err != nil {
			fmt.Printf("error compress: %v", err)
			return
		}
		fmt.Printf("%v: %v, gzipped: %7v, compressed size: %.1f%%\n",
			c.caseName, len(c.toCompress), len(gzipped), 100*float64(len(gzipped))/float64(len(c.toCompress)))
	}
}

func genRandomString(length int) string {
	charList := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	builder := strings.Builder{}
	builder.Grow(length)
	for i := 0; i < length; i++ {
		builder.WriteRune(charList[rand.Intn(len(charList))])
	}
	return builder.String()
}
