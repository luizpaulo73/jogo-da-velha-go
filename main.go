package main

import (
	"fmt"

	"jogo_da_velha/modos"
)

func main() {
	modoSelecionado := menu()
	switch modoSelecionado {
	case 1:
		modos.JogadorVsJogador()
	case 2:
		modos.JogadorVsMaquinaFacil()
	case 3:
		modos.JogadorVsMaquinaDificil()
	}
}

func menu() int {
	modoJogo := 0
	fmt.Println("Selecione o modo de jogo:")
	
	for modoJogo < 1 || modoJogo > 3 {
		fmt.Println("Digite: \n1 - para Jogador x Jogador \n" +
                             "2 - para Jogador x Máquina - nível fácil \n" +
                             "3 - para Jogador x Máquina - nível difícil")
		fmt.Scanln(&modoJogo)
	}
	return modoJogo
}