package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Uso: conversor <arquivo de entrada> <formato de saída>")
        fmt.Println("Exemplo: conversor video.mp4 avi")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    outputFormat := os.Args[2]
    outputFile := "output." + outputFormat

    if _, err := os.Stat(inputFile); os.IsNotExist(err) {
        fmt.Printf("Erro: O arquivo '%s' não foi encontrado!\n", inputFile)
        os.Exit(1)
    }

    fmt.Printf("Iniciando conversão de '%s' para '%s'...\n", inputFile, outputFile)
    cmd := exec.Command("ffmpeg", "-i", inputFile, "-preset", "fast", "-c:v", "libx264", "-c:a", "aac", outputFile)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("Erro na conversão do arquivo.")
        os.Exit(1)
    }

    fmt.Printf("Conversão concluída com sucesso! Arquivo gerado: %s\n", outputFile)
}
