package main

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func XPrintRow(row *xlsx.Row) {
	cellStrs := []string{}
	for _, cell := range row.Cells {
		cellS := cell.String()
		cellS = fmt.Sprintf("%12v", cellS)
		cellStrs = append(cellStrs, cellS)
	}
	result := strings.Join(cellStrs, " ")
	fmt.Println(result)
}

func main() {
	xlFile, err := xlsx.OpenFile("excel_read_write/people.xlsx")
	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		for _, sheet := range xlFile.Sheets {
			for i, row := range sheet.Rows {
				if i == 0 {
					fmt.Println("Fields: ")
					XPrintRow(row)
				} else {
					XPrintRow(row)
					for j, cell := range row.Cells {
						fmt.Println("cell ", j, cell.String())
					}
				}
			}
		}
	}
}
