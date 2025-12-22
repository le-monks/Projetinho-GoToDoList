package pipeline

import (
	"fmt"
	"os"
)

// Extract lê o arquivo do disco
func Extract(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Caminho de arquivo inválido")
	}
	return data, err
}
