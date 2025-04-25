package configuracao

import "fmt"

func configurarJogo() (string, string) {
	var simbolo int

	for simbolo < 1 || simbolo > 2 {
		fmt.Println("Escolha o símbolo do jogador 1: \n1 - X \n2 - O")
		fmt.Println("Digite o número correspondente ao símbolo desejado:")
		fmt.Scanln(&simbolo)
	}
	
	jogardor1, jogador2 := escolhaSimbolo(simbolo)

	if simbolo == 1 {
		fmt.Println("Jogador 1 escolheu X e Jogador 2 escolheu O")
	} else {
		fmt.Println("Jogador 1 escolheu O e Jogador 2 escolheu X")
	}

	return jogardor1, jogador2
}

func escolhaSimbolo(simbolo int) (string, string) {
	if (simbolo == 1) {
		return "X", "O"
	} else {
		return "O", "X"
	}

}

func EstruturaBasica() (string, string, int, int, string) {
	jogador1, jogador2 := configurarJogo()
	pontoJogador1 := 0
	pontoJogador2 := 0
	vezJogador := jogador1

	return jogador1, jogador2, pontoJogador1, pontoJogador2, vezJogador
}

func InicializarTabuleiro() map[string][]string {
	tabuleiro := map[string] []string {
		"linha1" : {" ", " ", " " },
		"linha2" : {" ", " ", " " },
		"linha3" : {" ", " ", " " },
	}
	return tabuleiro
}

func ImprimirTabuleiro(matriz map[string][]string) {
	linha1 := matriz["linha1"][0] + " | " + matriz["linha1"][1] + " | " + matriz["linha1"][2]
	linha2 := matriz["linha2"][0] + " | " + matriz["linha2"][1] + " | " + matriz["linha2"][2]
	linha3 := matriz["linha3"][0] + " | " + matriz["linha3"][1] + " | " + matriz["linha3"][2]

	fmt.Println("\n" + linha1 + "\n---------\n" + linha2 + "\n---------\n" + linha3)
}