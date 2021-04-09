package main

import (
	"log"
	"os"
	"reflect"
	"time"

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
		Name      string    `xlsx:"0" json:"Name"`
		WeightKgs float64   `xlsx:"1" json:"WeightKgs"`
		IsMale    bool      `xlsx:"2" json:"IsMale"`
		CreatedAt time.Time `xlsx:"3" json:"CreatedAt"` // not support RFC3339
		Nested    struct {
			Nested0 string
			Nested1 string
		} // will be ignore by xlsx_Row_WriteStruct
	}
	read := make([]Person, 0)
	xlFile.Sheets[0].ForEachRow(func(row *xlsx.Row) error {
		if row.GetCoordinate() == 0 { // skip rows that are field name
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
	log.Printf("read from people_read.xlsx: %#v", read)

	// write

	rowsToWrite := []Person{
		{Name: "Không còn không còn ai", WeightKgs: 11, IsMale: true, CreatedAt: time.Now()},
		{Name: "Ta trôi trong cuộc đời", WeightKgs: 22.2, IsMale: true, CreatedAt: time.Now()},
		{Name: "Không chờ không chờ ai", WeightKgs: 33333333333, IsMale: false, CreatedAt: time.Now()},
	}
	newFile := xlsx.NewFile() // just an object, need to write to a real file later
	sheet0, err := newFile.AddSheet("Sheet0")
	if err != nil {
		log.Fatal(err)
	}
	row0 := sheet0.AddRow()
	row0.WriteSlice(getStructTags(rowsToWrite[0], "json"), -1)
	row0.ForEachCell(func(cell *xlsx.Cell) error {
		cell.GetStyle().Font.Bold = true
		return nil
	})
	for _, row := range rowsToWrite {
		rowX := sheet0.AddRow()
		rowX.WriteStruct(&row, -1)
	}
	sheet0.SetColWidth(1, 1, 30)
	sheet0.SetColWidth(2, 2, 15)
	sheet0.SetColWidth(4, 4, 20)

	output, err := os.OpenFile("people_write.xlsx",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = newFile.Write(output)
	if err != nil {
		log.Fatal(err)
	}
	err = output.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote to people_write.xlsx")
}

func getStructTags(obj interface{}, tagToGet string) []string {
	tags := make([]string, 0)
	val := reflect.ValueOf(obj)
	for i := 0; i < val.Type().NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get(tagToGet)
		tags = append(tags, tag)
	}
	return tags
}
