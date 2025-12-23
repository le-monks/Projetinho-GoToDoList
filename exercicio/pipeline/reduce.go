package pipeline

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type ProductBase struct {
	XMLName      xml.Name `xml:""`
	ID           string   `xml:"id"`
	Title        string   `xml:"title,omitempty"`
	Description  string   `xml:"description,omitempty"`
	Link         string   `xml:"link,omitempty"`
	ImageLink    string   `xml:"image_link,omitempty"`
	Price        string   `xml:"price,omitempty"`
	Availability string   `xml:"availability,omitempty"`
	Brand        string   `xml:"brand,omitempty"`
}
type ProductStandard struct {
	XMLName      xml.Name `xml:""`
	ID           string   `xml:"id"`
	Title        string   `xml:"title,omitempty"`
	Description  string   `xml:"description,omitempty"`
	Link         string   `xml:"link,omitempty"`
	ImageLink    string   `xml:"image_link,omitempty"`
	Price        string   `xml:"price,omitempty"`
	Availability string   `xml:"availability,omitempty"`
	Brand        string   `xml:"brand,omitempty"`
}

// Reduce converte os dados brutos para o formato padrão
func Reduce(data []byte) ([]ProductStandard, error) {
	// TODO: Fazer unmarshal e normalização

	p := []ProductBase{}
	out := []ProductStandard{} 

	err := xml.Unmarshal(data, &p) // p deve ser um ponteiro para a struct (ex.: &Product{}). Se passar uma struct por valor, Unmarshal não conseguirá preencher os campos corretamente e pode retornar erro.

	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer unmarshal: %v", err)
	}

	for count := 0; count < len(p); count++ {
	
		out[count].ID = p[count].ID
		out[count].Title = p[count].Title
		out[count].Description = p[count].Description
		out[count].Link = p[count].Link
		out[count].ImageLink = p[count].ImageLink
		out[count].Price = FormatPriceBRL(p[count].Price)

		out[count].Availability = FormatStockStatus(strings.ToLower(p[count].Availability))
		out[count].Brand = p[count].Brand


	}


	return nil, nil
}


func FormatStockStatus(status string) string{

	// Format stock status to standard values
	// Standard values: "in-stock", "out-of-stock", "pre-order"

	// in-stock
	if status == "in stock" || status == "available" || status == "yes" || status == "em estoque" {
		return "in-stock"
	}

	// out-of-stock
	if status == "out of stock" || status == "unavailable" || status == "no" || status == "fora de estoque" {
		return "out-of-stock"
	}

	// pre-order
	if status == "pre order" || status == "pre-order" || status == "pré-venda" {
		return "pre-order"
	}

	return "unknown"
}

func FormatPriceBRL(price string) string {
	// Format price to BRL standard
	// Example: "R$ 1234.56" -> "1234,56 BRL"

	// Replace "." with ","
	s := strings.Replace(price, ".", ",", 1)

	// Remove R$ if it exists
	s = strings.Replace(s, "R$", "", 1)

	// Trim spaces
	s = strings.TrimSpace(s)

	// Add BRL if not present
	if !strings.HasSuffix(s, " BRL") {
		s = s + " BRL"
	}

	return s
}