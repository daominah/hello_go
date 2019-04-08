package main

import (
	"fmt"
)

type Human struct {
	Name     string
	BirthDay string
}

type Coder struct {
	Lang string
	Human
}

func (b *Human) GetName() string {
	return b.Name
}

func (c *Coder) GetName() string {
	return "Coder's method GetName: " + c.Name
}

func main() {
	h := Human{Name: "HumanName", BirthDay: "HumanBirthDay1992-08-20"}
	c := Coder{Human: h, Lang:"CoderLang go"}
	h.Name = "HumanNameChanged"
	//c := Coder{Human: Human{Name: "c0", BirthDay: "2018"}, Lang: "go"}  // right syntax
	//c := Coder{Name: "c0", BirthDay: "2018", Lang: "go"}  // wrong syntax
	fmt.Println("h.GetName",h.GetName())
	fmt.Println("c.GetName",c.GetName())
	fmt.Println("Attrs:", c.Name, c.BirthDay, c.Lang)
}
