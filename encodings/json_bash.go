package main

import (
	"encoding/json"
	"strings"

	"github.com/mywrap/log"
)

func main() {
	const SingleQuote = `\u0027`
	data := map[string]interface{}{
		"Key0": "ass hole's value",
	}

	originDumped, _ := json.Marshal(data)
	dumped := strings.ReplaceAll(string(originDumped), "'", SingleQuote)
	log.Printf("%v", dumped)

	var loaded map[string]interface{}
	json.Unmarshal([]byte(dumped), &loaded)
	log.Printf("%v", loaded["Key0"])
}
