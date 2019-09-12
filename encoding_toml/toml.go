package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	errors2 "github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strings"
)

func tomlToJson(tomled string) (string, error) {
	var obj interface{}
	_, err := toml.Decode(tomled, &obj)
	if err != nil {
		return "", errors2.Wrap(err, "cannot toml decode")
	}

	bs, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return "", errors2.Wrap(err, "cannot json encode")
	}
	jsoned := string(bs)

	return jsoned, nil
}

func jsonToToml(jsoned string) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(jsoned), &obj)
	if err != nil {
		return "", errors2.Wrap(err, "cannot json decode")
	}

	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	encoder.Indent = strings.Repeat(" ", 4)
	err = encoder.Encode(obj)
	if err != nil {
		return "", errors2.Wrap(err, "cannot toml encode")
	}
	tomled := buf.String()

	return tomled, nil
}

func main() {
	// try jsonToToml
	{
		jsoned := `{
		"field1": "value1",
		"field2": 2.2,
		"field3": true,
		"field3": [
			"value31",
			"value32"
		],
		"field4": {
			"field41": 41,
			"field42": [421, 422, 423]
		},
		"field5": [
			{"k1": 1, "k2": 2},
			{"k3": "v3", "k4": 4}
		]
	}`
		tomled, err := jsonToToml(jsoned)
		log.SetFlags(log.Ltime | log.Lshortfile)
		if err != nil {
			log.Fatal(err)
		}
		_, _ = fmt.Println, tomled
		fmt.Println("tomled:\n", tomled)
	}

	// try tomlToJson
	{
		tomled := `
			# lang_pack_en.toml
			
			langCode = "en"
			version = 14897
			
			[[Strings]]
				value = "Your code"
				key = "lng_code_ph"
				
				[[Strings]]
				value = "We've sent a code [b]via Telegram[/b]\nto your other devices. Please enter it below."
				key = "lng_code_telegram"

			[[StringPluralizeds]]
				fewValue = ""
				oneValue = "{count} minute"
				key = "lng_signin_reset_minutes"
		`
		jsoned, err := tomlToJson(tomled)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("jsoned:\n", jsoned)
	}

	bs, err := ioutil.ReadFile("/home/tungdt/go/src/github.com/daominah/translate/data_kingtalk/lang_pack_en.toml")
	if err != nil {
		log.Fatal(err)
	}
	j, err := tomlToJson(string(bs))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ngon", j)
}
