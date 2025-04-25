package acoes

import (
	"fmt"
	"jogo_da_velha/utils"
)

func RealizarJogada(numeroUtilizado []int) int {

	jogada := 0
	
	for jogada < 1 || jogada > 9 || utils.JaUtilizado(jogada, numeroUtilizado) {
		fmt.Println("Faça a sua jogada 1 - 9:")
		fmt.Scanln(&jogada)
	}

	return jogada
}

func MapearJogada(jogada int, tabuleiro map[string][]string, vezJogador string) {
	jogada -= 1
	if jogada >= 0 && jogada <= 2 {
		tabuleiro["linha1"][jogada] = vezJogador
	} else if jogada >= 3 && jogada <= 5 {
		tabuleiro["linha2"][jogada - 3] = vezJogador
	} else if jogada >= 6 && jogada <= 9 {
		tabuleiro["linha3"][jogada - 6] = vezJogador
	}
}

func VerificarVencedor(tabuleiro map[string][]string) string {
	// verificar vitoria horizontal
	for i := 1; i <= 3; i++ {
		linha := fmt.Sprintf("linha%d", i)
		if tabuleiro[linha][0] == tabuleiro[linha][1] && tabuleiro[linha][1] == tabuleiro[linha][2] && tabuleiro[linha][0] != " " {
			return tabuleiro[linha][0]
		}
	}

	// verificar vitoria vertical
	for j := 0; j <= 2; j++ {
		if tabuleiro["linha1"][j] == tabuleiro["linha2"][j] && tabuleiro["linha2"][j] == tabuleiro["linha3"][j] && tabuleiro["linha1"][j] != " " {
			return tabuleiro["linha1"][j]
		}
	}

	//verificar vitoria diagonal
	if tabuleiro["linha1"][0] == tabuleiro["linha2"][1] && tabuleiro["linha2"][1] == tabuleiro["linha3"][2] && tabuleiro["linha2"][1] != " " {
		return tabuleiro["linha2"][1]
	}
	if tabuleiro["linha1"][2] == tabuleiro["linha2"][1] && tabuleiro["linha2"][1] == tabuleiro["linha3"][0] && tabuleiro["linha2"][1] != " " {
		return tabuleiro["linha2"][1]
	}

	return ""
}