package main

import (
	"log"

	xlsx "github.com/tealeg/xlsx/v3"
)

func main() {
	// read

	xlFile, err := xlsx.OpenFile("people_read.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	if len(xlFile.Sheets) == 0 {
		log.Fatal("error the file has no sheets")
	}
	type Person struct {
		Name      string  `xlsx:"0"`
		WeightKgs float64 `xlsx:"1"`
		IsMale    bool    `xlsx:"2"`
	}
	read := make([]Person, 0)
	xlFile.Sheets[0].ForEachRow(func(row *xlsx.Row) error {
		if row.GetCoordinate() == 0 {
			return nil
		}
		var person Person
		err := row.ReadStruct(&person)
		if err != nil {
			log.Printf("error read row %v: %v", row.GetCoordinate(), err)
		}
		read = append(read, person)
		return nil
	})
	log.Printf("read: %#v", read)
}
