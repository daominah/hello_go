package main

import (
	"encoding/xml"
	"log"
)

type Lala struct {
	CuoiXinh bool
	NgoNgo   string
	Emo      []int
}

func main() {
	lan := Lala{CuoiXinh: true, NgoNgo: "y", Emo: []int{2, 3, 4}}
	temp, err := xml.MarshalIndent(lan, "", "    ")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("temp:\n%s", temp)

	//s := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	//		<soap:Body>
	//			<soap:Fault>
	//				<faultcode>soap:Server</faultcode>
	//				<faultstring>Fault occurred while processing.</faultstring>
	//					<detail>
	//						<ns1:ExceptionName xmlns:ns1="http://ws.token.authentication.fo.fss.com/"/>
	//					</detail>
	//			</soap:Fault>
	//		</soap:Body>
	//	</soap:Envelope>`
	//_ = s

	var v Lala
	err = xml.Unmarshal(temp, &v)
	if err != nil {
		log.Fatal("cannot unmarshal", err)
	}
	log.Printf("v: %#v", v)
}
