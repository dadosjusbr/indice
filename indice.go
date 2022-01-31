package indice

import (
	"github.com/dadosjusbr/proto/coleta"
)

type Score struct {
	Score             float64
	CompletenessScore float64
	EasinessScore     float64
}

func calcCriteria(criteria bool, value float64) float64 {
	if criteria {
		return value
	}
	return 0
}

func calcStringCriteria(criteria string, values map[string]float64) float64 {
	for k := range values {
		if criteria == k {
			return values[k]
		}
	}
	return 0
}

func calcCompletenessScore(meta coleta.Metadados) float64 {
	var score float64 = 0
	var options = map[string]float64{"SUMARIZADO": 0.5, "DETALHADO": 1}

	score = score + calcCriteria(meta.TemLotacao, 1)
	score = score + calcCriteria(meta.TemCargo, 1)
	score = score + calcCriteria(meta.TemMatricula, 1)
	score = score + calcStringCriteria(meta.ReceitaBase.String(), options)
	score = score + calcStringCriteria(meta.OutrasReceitas.String(), options)
	score = score + calcStringCriteria(meta.Despesas.String(), options)

	return score / 6
}

func calcEasinessScore(meta coleta.Metadados) float64 {
	var score float64 = 0
	var options = map[string]float64{
		"ACESSO_DIRETO":          1,
		"AMIGAVEL_PARA_RASPAGEM": 0.5,
		"RASPAGEM_DIFICULTADA":   0.25}

	score = score + calcCriteria(meta.NaoRequerLogin, 1)
	score = score + calcCriteria(meta.NaoRequerCaptcha, 1)
	score = score + calcStringCriteria(meta.Acesso.String(), options)
	score = score + calcCriteria(meta.FormatoConsistente, 1)
	score = score + calcCriteria(meta.EstritamenteTabular, 1)

	return score / 5
}

func calcScore(meta coleta.Metadados) float64 {
	var score = 0.0
	var completeness = calcCompletenessScore(meta)
	var easiness = calcEasinessScore(meta)
	score = (completeness + easiness) / 2

	return score
}

func CalcScore(meta coleta.Metadados) Score {
	var score = 0.0
	var completeness = calcCompletenessScore(meta)
	var easiness = calcEasinessScore(meta)
	score = (completeness + easiness) / 2

	return Score{
		Score:             score,
		CompletenessScore: completeness,
		EasinessScore:     easiness,
	}
}
