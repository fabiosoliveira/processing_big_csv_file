package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Abrir o arquivo para leitura
	f, err := os.Open("train.csv")
	if err != nil {
		panic(err)
	}

	// Ler o conteúdo do arquivo, ignorando a primeira linha
	scanner := bufio.NewScanner(f)
	var lines []string
	isFirstLine := true
	for scanner.Scan() {
		if isFirstLine {
			isFirstLine = false
			continue
		}
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	f.Close()

	// Reabrir o arquivo no modo de escrita
	f, err = os.Create("train.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Escrever o conteúdo lido 10 vezes no arquivo
	writer := bufio.NewWriter(f)
	for i := 0; i < 217; i++ {
		for _, line := range lines {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	writer.Flush()

	fmt.Println("Arquivo processado com sucesso.")
}
