package pipeline

import (
	"io"
	"os"
)

// Extract lê o arquivo do disco
func Extract(filename string) ([]byte, error) {
	// TODO: Implementar leitura de arquivo
	file, err := os.Open(filename);

	if err != nil{
		return nil, err; 
	}

	defer file.Close()

	data, err := io.ReadAll(file);

	if err != nil{
		return nil, err; 
	}

	return data, nil;
}
