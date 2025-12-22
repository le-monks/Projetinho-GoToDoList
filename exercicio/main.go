package main

import (
	"exercicio/pipeline"
	"fmt"
	"log"
	"os"
	//	"log"
	//	"exercicio/pipeline"
)

func main() {
	fmt.Println("🚀 Iniciando Pipeline Local...")
	filepath := os.Getenv("FEED_BRUTO_PATH")
	if filepath == "" {
		log.Fatal(fmt.Errorf("Variável de ambiente não encontrada ou nula"))
	}
	// 1. Extract
	// TODO: Implementar chamada
	fmt.Println("1. Lendo arquivo...")

	rawData, err := pipeline.Extract(filepath)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(string(rawData))

	// 2. Reduce
	// TODO: Implementar chamada
	//fmt.Println("2. Normalizando dados...")
	// products, err := pipeline.Reduce(rawData)

	// 3. FFT
	// TODO: Implementar chamada
	// fmt.Println("3. Aplicando regras de negócio (Google)...")
	// googleProducts := pipeline.FFT(products)

	// 4. Load
	// TODO: Implementar chamada
	// fmt.Println("4. Gerando arquivo final...")
	// err := pipeline.Load(googleProducts, "feed_google.xml")

	// fmt.Println("✅ Processo finalizado com sucesso!")
}
