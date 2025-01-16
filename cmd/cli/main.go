package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Statistic struct {
	TotalMale   int
	TotalFemale int
	FaixaEtaria map[string]int
}

func (s *Statistic) AddMale() {
	s.TotalMale++
}

func (s *Statistic) AddFemale() {
	s.TotalFemale++
}

func (s *Statistic) AddFaixaEtaria(age int) {
	switch {
	case age <= 18:
		s.FaixaEtaria["0-18"]++
	case age <= 30:
		s.FaixaEtaria["19-30"]++
	case age <= 50:
		s.FaixaEtaria["31-50"]++
	case age <= 70:
		s.FaixaEtaria["51-70"]++
	default:
		s.FaixaEtaria["71+"]++
	}
}

func (s Statistic) String() string {
	return fmt.Sprintf(`
Total de Pacientes por sexo:
  Total Male:   %d
  Total Female: %d

Total por Faixas Etárias:
  0-18:  %d
  19-30: %d
  31-50: %d
  51-70: %d
  71+:   %d
	`, s.TotalMale, s.TotalFemale, s.FaixaEtaria["0-18"], s.FaixaEtaria["19-30"], s.FaixaEtaria["31-50"], s.FaixaEtaria["51-70"], s.FaixaEtaria["71+"])
}

// O arquivo train.csv foi baixado do Kaggle: https://www.kaggle.com/datasets/ashery/chexpert?resource=download

func main() {
	f, err := os.Open("train.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	statistics := Statistic{
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
			statistics.AddMale()
		} else {
			statistics.AddFemale()
		}

		age, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error converting age to int:", err)
			return
		}

		statistics.AddFaixaEtaria(age)

		if isPrefixed {
			break
		}

	}

	fmt.Println(statistics)

	fmt.Println("\n\nMemory Stats:")
	printMemStats()
}

func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
