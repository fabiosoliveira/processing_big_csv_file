package demografic

import "fmt"

type DemographicData struct {
	TotalMale   int
	TotalFemale int
	FaixaEtaria map[string]int
}

func NewDemographicData() *DemographicData {
	return &DemographicData{
		FaixaEtaria: make(map[string]int),
	}
}

func (s *DemographicData) AddMale() {
	s.TotalMale++
}

func (s *DemographicData) AddFemale() {
	s.TotalFemale++
}

func (s *DemographicData) AddFaixaEtaria(age int) {
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

func (s DemographicData) String() string {
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
