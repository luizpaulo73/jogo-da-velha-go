package modos

import (
	"fmt"

	"jogo_da_velha/configuracao"
	"jogo_da_velha/acoes"
	"jogo_da_velha/jogadaBot"
)

func JogadorVsMaquinaFacil() {
	jogador1, jogador2, pontoJogador1, pontoJogador2, vezJogador := configuracao.EstruturaBasica()

	for pontoJogador1 < 3 && pontoJogador2 < 3 {
		tabuleiro := configuracao.InicializarTabuleiro()
		turno := 1
		vencedor := ""
		numeroUtilizado := []int{}

		for turno <= 9 && vencedor == "" {
			configuracao.ImprimirTabuleiro(tabuleiro)
			fmt.Printf("Vez do jogador: %v", vezJogador)

			var jogada int
			if vezJogador == jogador1 {
				jogada = acoes.RealizarJogada(numeroUtilizado)
			} else {
				jogada = jogadaBot.JogadaNivelFacil(numeroUtilizado)
			}
			numeroUtilizado = append(numeroUtilizado, jogada)
			acoes.MapearJogada(jogada, tabuleiro, vezJogador)
			// verificar vencedor
			vencedor = acoes.VerificarVencedor(tabuleiro)
			if vencedor != "" {
				configuracao.ImprimirTabuleiro(tabuleiro)
				fmt.Printf("O jogador %v venceu!", vencedor)
				if vencedor == jogador1 {
					pontoJogador1++
				} else {
					pontoJogador2++
				}
				vezJogador = vencedor

			} else {
				if vezJogador == jogador1 {
					vezJogador = jogador2
				} else {
					vezJogador = jogador1
				}
			}
			turno++
		}
	}
	// imprime o vencedor e mensagem de vitoria
	var vencedorFinal string
	if pontoJogador1 == 3 {
		vencedorFinal = jogador1
	} else {
		vencedorFinal = jogador2
	}
	fmt.Printf("O jogador %v venceu o jogo! \nPor %d-%d", vencedorFinal, pontoJogador1, pontoJogador2)
}