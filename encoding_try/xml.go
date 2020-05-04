package main

import (
	"encoding/xml"
	"log"

	"encoding/json"
)

type ExHtml struct {
	XMLName xml.Name
	Lang    string `xml:"lang,attr"`
	NgApp   string `xml:"ng-app,attr"`
	Head    struct {
		Script struct {
			CharData string `xml:",chardata"`
		} `xml:"script"`
		Metas []struct {
			CharSet string `xml:"charset,attr,omitempty"`
			Name    string `xml:"name,attr,omitempty"`
			Content string `xml:"content,attr,omitempty"`
		} `xml:"meta"`
		Link struct {
			Name string `xml:"rel,attr"`
			Size string `xml:"sizes,attr"`
			Href string `xml:"href,attr"`
		} `xml:"link"`
	} `xml:"head"`
	Body struct {
		Div struct {
			A struct {
				Class    string `xml:"class,attr"`
				Href     string `xml:"href,attr"`
				CharData string `xml:",chardata"`
			} `xml:"a"`
		} `xml:"div"`
		Script struct {
			Src string `xml:"src,attr"`
		} `xml:"script"`
	} `xml:"body"`
}

func (h *ExHtml) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias ExHtml
	d.Strict = false
	d.AutoClose = xml.HTMLAutoClose
	return d.DecodeElement((*Alias)(h), &start)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	data := `<!doctype html>
<html lang="en" ng-app="tour">

<head>
    <script>function gtag() { dataLayer.push(arguments); }</script>
    <meta charset="utf-8">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <link rel="shortcut icon" sizes="196x196" href="/favicon.ico">
</head>

<body>
    <div class="bar top-bar">
        <a class="left logo" href="/list">A Tour of Go</a>
    </div>
    <script src="/script.js"></script>
</body>

</html>`

	var v ExHtml
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Fatal("cannot unmarshal", err)
	}

	beauty, _ := json.MarshalIndent(v, "", "    ")
	log.Printf("beauty: %s\n", beauty)

	myXMLDumped, _ := xml.MarshalIndent(v, "", "    ")
	log.Printf("myXMLDumped: %s\n", myXMLDumped)
}
