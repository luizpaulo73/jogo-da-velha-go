package jogadaBot

import (
	"fmt"
	"jogo_da_velha/utils"
	"math/rand"
	"time"
)

func JogadaNivelDificil(tabuleiro map[string][]string, jogador1 string, jogador2 string, numeroUtilizado []int) int {
	if tabuleiro["linha2"][1] == " " {
		return 5
	}

	jogadaBot := ganharOuBloquear(tabuleiro, "", jogador2)
	if jogadaBot != -1 && !utils.JaUtilizado(jogadaBot, numeroUtilizado) {
		return jogadaBot
	}

	jogadaBot = ganharOuBloquear(tabuleiro, jogador1, "")
	if jogadaBot != -1 && !utils.JaUtilizado(jogadaBot, numeroUtilizado) {
		return jogadaBot
	}

	jogadaBot = cantoVazio(tabuleiro)
	if jogadaBot != -1 && !utils.JaUtilizado(jogadaBot, numeroUtilizado) {
		return jogadaBot
	}

	jogadaBot = bordaVazia(numeroUtilizado)
	if jogadaBot != -1 && !utils.JaUtilizado(jogadaBot, numeroUtilizado) {
		return jogadaBot
	}

	return jogadaBot
}

func ganharOuBloquear(tabuleiro map[string][]string, jogador1 string, jogador2 string) int {
	var jogador string

	if jogador1 != "" {
		jogador = jogador1
	} else {
		jogador = jogador2
	}

	for i := 1; i <= 3; i++ {
		// Verifica as linhas
		linha := fmt.Sprintf("linha%d", i)
		if utils.CountValor(tabuleiro, linha, jogador) == 2 && utils.Contains(tabuleiro, linha) {
			return utils.Index(tabuleiro[linha], " ") + 1 + (i-1)*3
		}

		// Verifica as colunas
		coluna := []string{tabuleiro["linha1"][i-1], tabuleiro["linha2"][i-1], tabuleiro["linha3"][i-1]}
		if countSlice(coluna, jogador) == 2 && utils.ContainsTabuleiro(coluna, " ") {
			return utils.Index(coluna, " ")*3 + i
		}
	}

	// Verificar diagonais
	diagonal1 := []string{tabuleiro["linha1"][0], tabuleiro["linha2"][1], tabuleiro["linha3"][2]}
	if countSlice(diagonal1, jogador) == 2 && utils.ContainsTabuleiro(diagonal1, " ") {
		return utils.Index(diagonal1, " ")*4 + 1
	}

	diagonal2 := []string{tabuleiro["linha1"][2], tabuleiro["linha2"][1], tabuleiro["linha3"][0]}
	if countSlice(diagonal2, jogador) == 2 && utils.ContainsTabuleiro(diagonal2, " ") {
		return utils.Index(diagonal2, " ")*2 + 3
	}

	return -1
}

func countSlice(slice []string, valor string) int {
	count := 0
	for _, v := range slice {
		if v == valor {
			count++
		}
	}
	return count
}

func cantoVazio(tabuleiro map[string][]string) int {
	cantos := []int{1, 3, 7, 9}

	for _, canto := range(cantos) {
		if canto == 1 && tabuleiro["linha1"][0] == " " {
			return 1
		} else if canto == 3 && tabuleiro["linha1"][2] == " " {
			return 3
		} else if canto == 7 && tabuleiro["linha3"][0] == " " {
			return 7
		} else if canto == 9 && tabuleiro["linha3"][2] == " " {
			return 9
		}
	}
	return -1
}

func bordaVazia(numeroUtilizado []int) int {
	bordas := []int{2, 4, 6, 8}

	var jogada int

	for {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		jogada = r.Intn(4)

		if !utils.JaUtilizado(bordas[jogada], numeroUtilizado) {
			break
		}
	}

	return bordas[jogada]
}