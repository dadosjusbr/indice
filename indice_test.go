package indice

import (
	"testing"

	"github.com/dadosjusbr/proto/coleta"
	"github.com/stretchr/testify/assert"
)

func TestCalcCompletenessScore(t *testing.T) {
	data := []struct {
		Desc     string
		Input    coleta.Metadados
		Expected float64
	}{
		{"Sempre positivo", coleta.Metadados{
			TemMatricula:   true,
			TemLotacao:     true,
			TemCargo:       true,
			ReceitaBase:    coleta.Metadados_DETALHADO,
			OutrasReceitas: coleta.Metadados_DETALHADO,
			Despesas:       coleta.Metadados_DETALHADO,
		}, 1.0},
		{"Sempre negativo", coleta.Metadados{
			TemMatricula:   false,
			TemLotacao:     false,
			TemCargo:       false,
			ReceitaBase:    coleta.Metadados_AUSENCIA,
			OutrasReceitas: coleta.Metadados_AUSENCIA,
			Despesas:       coleta.Metadados_AUSENCIA,
		}, 0.0},
		{"CNJ-2020", coleta.Metadados{
			TemMatricula:   false,
			TemLotacao:     false,
			TemCargo:       false,
			ReceitaBase:    coleta.Metadados_DETALHADO,
			OutrasReceitas: coleta.Metadados_DETALHADO,
			Despesas:       coleta.Metadados_DETALHADO,
		}, 0.5},
		{"mppb-11-2021", coleta.Metadados{
			TemMatricula:   true,
			TemLotacao:     true,
			TemCargo:       true,
			ReceitaBase:    coleta.Metadados_DETALHADO,
			OutrasReceitas: coleta.Metadados_DETALHADO,
			Despesas:       coleta.Metadados_DETALHADO,
		}, 1},
		{"mpto-6-2019", coleta.Metadados{
			TemMatricula:   true,
			TemLotacao:     true,
			TemCargo:       true,
			ReceitaBase:    coleta.Metadados_DETALHADO,
			OutrasReceitas: coleta.Metadados_DETALHADO,
			Despesas:       coleta.Metadados_DETALHADO,
		}, 1},
	}

	for _, d := range data {
		t.Run(d.Desc, func(t *testing.T) {
			b := calcCompletenessScore(d.Input)
			assert.Equal(t, d.Expected, b)
		})
	}
}

func TestCalcEasinessScore(t *testing.T) {
	data := []struct {
		Desc     string
		Input    coleta.Metadados
		Expected float64
	}{
		{"Sempre positivo", coleta.Metadados{
			Acesso:              coleta.Metadados_ACESSO_DIRETO,
			FormatoConsistente:  true,
			EstritamenteTabular: true,
			FormatoAberto:       true,
		}, 1.0},
		{"Sempre negativo", coleta.Metadados{
			Acesso:              coleta.Metadados_NECESSITA_SIMULACAO_USUARIO,
			FormatoConsistente:  false,
			EstritamenteTabular: false,
			FormatoAberto:       false,
		}, 0.0},
		{"CNJ-2020", coleta.Metadados{
			Acesso:              coleta.Metadados_NECESSITA_SIMULACAO_USUARIO,
			FormatoConsistente:  true,
			EstritamenteTabular: true,
			FormatoAberto:       false,
		}, 0.5},
		{"mppb-11-2021", coleta.Metadados{
			Acesso:              coleta.Metadados_ACESSO_DIRETO,
			FormatoConsistente:  true,
			EstritamenteTabular: false,
			FormatoAberto:       true,
		}, 0.75},
		{"mpto-6-2019", coleta.Metadados{
			NaoRequerLogin:      true,
			NaoRequerCaptcha:    true,
			Acesso:              coleta.Metadados_RASPAGEM_DIFICULTADA,
			FormatoConsistente:  false,
			EstritamenteTabular: false,
			FormatoAberto:       false,
		}, 0.125},
	}

	for _, d := range data {
		t.Run(d.Desc, func(t *testing.T) {
			b := calcEasinessScore(d.Input)
			assert.Equal(t, d.Expected, b)
		})
	}
}

func TestCalcScore(t *testing.T) {
	data := []struct {
		Desc     string
		Input    coleta.Metadados
		Expected Score
	}{
		{"Sempre positivo", coleta.Metadados{
			TemMatricula:        true,
			TemLotacao:          true,
			TemCargo:            true,
			ReceitaBase:         coleta.Metadados_DETALHADO,
			OutrasReceitas:      coleta.Metadados_DETALHADO,
			Despesas:            coleta.Metadados_DETALHADO,
			Acesso:              coleta.Metadados_ACESSO_DIRETO,
			FormatoConsistente:  true,
			EstritamenteTabular: true,
			FormatoAberto:       true,
		}, Score{1, 1, 1}},
		{"Sempre negativo", coleta.Metadados{
			TemMatricula:        false,
			TemLotacao:          false,
			TemCargo:            false,
			ReceitaBase:         coleta.Metadados_AUSENCIA,
			OutrasReceitas:      coleta.Metadados_AUSENCIA,
			Despesas:            coleta.Metadados_AUSENCIA,
			Acesso:              coleta.Metadados_NECESSITA_SIMULACAO_USUARIO,
			FormatoConsistente:  false,
			EstritamenteTabular: false,
			FormatoAberto:       false,
		}, Score{0, 0, 0}},
		{"CNJ-2020", coleta.Metadados{
			TemMatricula:        false,
			TemLotacao:          false,
			TemCargo:            false,
			ReceitaBase:         coleta.Metadados_DETALHADO,
			OutrasReceitas:      coleta.Metadados_DETALHADO,
			Despesas:            coleta.Metadados_DETALHADO,
			Acesso:              coleta.Metadados_NECESSITA_SIMULACAO_USUARIO,
			FormatoConsistente:  true,
			EstritamenteTabular: true,
			FormatoAberto:       false,
		}, Score{0.5, 0.5, 0.5}},
		{"mppb-11-2021", coleta.Metadados{
			TemMatricula:        true,
			TemLotacao:          true,
			TemCargo:            true,
			ReceitaBase:         coleta.Metadados_DETALHADO,
			OutrasReceitas:      coleta.Metadados_DETALHADO,
			Despesas:            coleta.Metadados_DETALHADO,
			Acesso:              coleta.Metadados_ACESSO_DIRETO,
			FormatoConsistente:  true,
			EstritamenteTabular: false,
			FormatoAberto:       true,
		}, Score{0.85, 1, 0.75}},
		{"mpto-6-2019", coleta.Metadados{
			TemMatricula:        true,
			TemLotacao:          true,
			TemCargo:            true,
			ReceitaBase:         coleta.Metadados_DETALHADO,
			OutrasReceitas:      coleta.Metadados_DETALHADO,
			Despesas:            coleta.Metadados_DETALHADO,
			Acesso:              coleta.Metadados_RASPAGEM_DIFICULTADA,
			FormatoConsistente:  false,
			EstritamenteTabular: false,
			FormatoAberto:       false,
		}, Score{0.22, 1, 0.125}},
	}

	for _, d := range data {
		t.Run(d.Desc, func(t *testing.T) {
			b := CalcScore(d.Input)
			assert.InDelta(t, d.Expected.Score, b.Score, 0.01)
			assert.InDelta(t, d.Expected.CompletenessScore, b.CompletenessScore, 0.01)
			assert.InDelta(t, d.Expected.EasinessScore, b.EasinessScore, 0.01)
		})
	}
}
