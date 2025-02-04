package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mywrap/textproc"
)

func main() {
	log.SetFlags(log.Lshortfile)

	inputFilePath := `C:/Users/tungd/OneDrive/Desktop/a.html`

	data, err := os.ReadFile(inputFilePath)
	if err != nil {
		log.Fatalf("error os.ReadFile: %v", err)
	}
	root := textproc.HTMLParseToNode(data)

	rows, err := textproc.HTMLXPath(root, `//*[@role="row"]`)
	if err != nil {
		log.Fatalf("error HTMLXPath: %v", err)
	}
	log.Printf("len rows in HTML: %v", len(rows))

	var items []CartItem
	for _, row := range rows {
		cells, err := textproc.HTMLXPath(row, `//*[@role="cell"]`)
		if err != nil {
			log.Fatalf("error HTMLXPath: %v", err)
		}
		if len(cells) < 3 {
			continue
		}

		var item CartItem
		imgNodes, _ := textproc.HTMLXPath(cells[0], `//img`)
		if len(imgNodes) > 0 {
			item.ImageURL = textproc.HTMLGetImgSrc("", imgNodes[0])
		}
		spans, _ := textproc.HTMLXPath(cells[0], `//span`)
		if len(spans) > 0 {
			quantityStr := textproc.HTMLGetText(spans[len(spans)-1])
			item.Quantity, _ = strconv.Atoi(quantityStr)
		}

		spansC1, _ := textproc.HTMLXPath(cells[1], `//span`)
		if len(spans) < 1 {
			log.Printf("unexpected len spansC1: %v, html: %v", len(spansC1), textproc.HTMLRender(cells[1]))
			continue
		}
		numberNameRarity := textproc.HTMLGetText(spansC1[0])
		parts := strings.Fields(numberNameRarity)
		if len(parts) >= 3 {
			item.CardNumber = parts[0]
			item.Rarity = mapAbbreviation(parts[len(parts)-1])
			item.CardName = strings.Join(parts[1:len(parts)-1], " ")
		}
		item.PriceAll = parsePrice(textproc.HTMLGetText(cells[2]))
		item.PriceOne = item.PriceAll / float64(item.Quantity)

		items = append(items, item)
	}

	// write parsed items to CSV
	file, err := os.Create("tcgcorner_order.csv")
	if err != nil {
		log.Fatalf("error creating file: %v", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	header := []string{"CardNumber", "CardName", "Rarity", "PriceOne", "PriceAll", "Quantity", "ImageURL"}
	if err := writer.Write(header); err != nil {
		log.Fatalf("error writing header to csv: %v", err)
	}
	for _, item := range items {
		record := []string{
			item.CardNumber,
			item.CardName,
			item.Rarity,
			strconv.FormatFloat(item.PriceOne, 'f', 0, 64),
			strconv.FormatFloat(item.PriceAll, 'f', 0, 64),
			strconv.Itoa(item.Quantity),
			// item.ImageURL,
		}
		if err := writer.Write(record); err != nil {
			log.Fatalf("error writing record to csv: %v", err)
		}
		writer.Flush()
	}
	log.Printf("write %v items to csv", len(items))
}

type CartItem struct {
	CardNumber string // e.g. "RA03-EN079"
	CardName   string
	Rarity     string
	PriceOne   float64
	PriceAll   float64 // should equal to PriceOne * Quantity
	Quantity   int
	ImageURL   string
}

func mapAbbreviation(abbreviation string) string {
	abbreviation = strings.TrimPrefix(abbreviation, "(")
	abbreviation = strings.TrimSuffix(abbreviation, ")")
	switch abbreviation {
	case "N":
		return "Common"
	case "R":
		return "Rare"
	case "SR":
		return "Super Rare"
	case "UR":
		return "Ultra Rare"
	case "SE", "SER":
		return "Secret Rare"
	case "UL":
		return "Ultimate Rare"
	case "CR":
		return "Collector Rare"
	case "ES", "EXSER":
		return "Extra Secret Rare"
	case "PS":
		return "Platinum Secret Rare"
	case "QCSE", "QCSR":
		return "Quarter Century Secret Rare"
	case "HR":
		return "Holographic Rare"
	}
	return abbreviation
}

func parsePrice(priceStr string) float64 {
	priceStr = strings.Map(func(r rune) rune { // remove all non-digit characters
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, priceStr)
	price, _ := strconv.ParseFloat(priceStr, 64)
	return price
}
