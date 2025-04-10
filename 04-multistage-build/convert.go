package main

import (
	"fmt"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <arquivo de entrada> <formato de saída>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFormat := os.Args[2]
	outputFile := fmt.Sprintf("output.%s", outputFormat)

	fmt.Printf("Convertendo %s para %s...\n", inputFile, outputFile)

	err := ffmpeg.Input(inputFile).
		Output(outputFile).
		OverWriteOutput().
		Run()

	if err != nil {
		fmt.Printf("Erro na conversão: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Conversão finalizada! Arquivo gerado: %s\n", outputFile)
}
