package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fabiosoliveira/processing_big_csv_file/pkg/demografic"
	"github.com/fabiosoliveira/processing_big_csv_file/pkg/stats"
)

// O arquivo train.csv foi baixado do Kaggle: https://www.kaggle.com/datasets/ashery/chexpert?resource=download

func main() {
	timeNow := time.Now()

	f, err := os.Open("train.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	demograficData := demografic.DemographicData{
		FaixaEtaria: make(map[string]int),
	}

	r4 := bufio.NewReader(f)
	r4.ReadLine()

	for {
		line, isPrefixed, err := r4.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading file:", err)
			return
		}

		parts := strings.Split(string(line), ",")

		sex := parts[1]

		if sex == "Male" {
			demograficData.AddMale()
		} else {
			demograficData.AddFemale()
		}

		age, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error converting age to int:", err)
			return
		}

		demograficData.AddFaixaEtaria(age)

		if isPrefixed {
			break
		}
	}

	fmt.Println(demograficData)

	fmt.Println("\n\nMemory Stats:")
	stats.PrintMemStats()

	fmt.Println("Time elapsed: ", time.Since(timeNow))
}
