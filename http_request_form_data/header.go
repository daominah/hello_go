package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func DumpHeader(header http.Header) string {
	kvs := make([]string,0)
	for key, values := range header {
		if len(values) > 0 {
			kvs = append(kvs, fmt.Sprintf("%v: %v", key, values[0]))
		}
	}
	sort.Strings(kvs)
	return strings.Join(kvs, "\n")
}
