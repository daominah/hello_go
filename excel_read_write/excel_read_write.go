package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"sort"

	"github.com/mywrap/log"
	"github.com/mywrap/textproc"
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
	xlFile, err := xlsx.OpenFile("/home/tungdt/Desktop/client_name_sub.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	if len(xlFile.Sheets) == 0 {
		log.Fatal("no sheets")
	}

	accs := make(map[string]string)
	subAccs := make(map[string]string)

	sheet0 := xlFile.Sheets[0]
	for i, row := range sheet0.Rows {
		if i == 0 {
			fmt.Println("Fields: ")
			XPrintRow(row)
		} else {
			if len(row.Cells) != 3 {
				continue
			}
			acc, name, subAcc := row.Cells[0].String(), row.Cells[1].String(), row.Cells[2].String()
			acc, name, subAcc = strings.TrimSpace(acc), strings.TrimSpace(name), strings.TrimSpace(subAcc)
			name = textproc.RemoveVietnamDiacritic(name)
			name = strings.ToUpper(name)
			name = strings.ReplaceAll(name, `"`, ``)
			accs[acc] = name
			subAccs[subAcc] = name
		}

		//if i == 10 {
		//	break
		//}
	}

	rows := make([][]string, 0)
	for k, v := range accs {
		rows = append(rows, []string{k, v})
	}
	for k, v := range subAccs {
		rows = append(rows, []string{k, v})
	}
	sort.Sort(SortByAcc(rows))
	log.Debugf("%#v", rows)

	f, _ := os.OpenFile("/home/tungdt/Desktop/names.csv",
		os.O_TRUNC|os.O_CREATE|os.O_RDWR,
		0644)
	w := csv.NewWriter(f)
	for _, record := range rows {
		if err := w.Write(record); err != nil {
			log.Fatal("error writing record to csv:", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatalf("error Flush: %v", err)
	}
	log.Printf("ok")
}

type SortByAcc [][]string

func (a SortByAcc) Len() int      { return len(a) }
func (a SortByAcc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByAcc) Less(i, j int) bool {
	if len(a[i]) < 1 || len(a[j]) < 1 {
		return false
	}
	return a[i][0] < a[j][0]
}
