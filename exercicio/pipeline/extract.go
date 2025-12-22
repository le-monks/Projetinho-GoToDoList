package pipeline

import (
	"fmt"
	"log"
	"os"
)

// Extract lê o arquivo do disco
func Extract(filename string) ([]byte, error) {
	if filename == "" {
		log.Fatal(fmt.Errorf("Variável de ambiente não encontrada ou nula"))
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Caminho de arquivo inválido")
	}
	return data, err
}
