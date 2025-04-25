package utils

func JaUtilizado(jogada int, numeroUtilizado []int) bool {
	for _, num := range numeroUtilizado {
		if num == jogada {
			return true
		}
	}
	return false
}

// Função para encontrar o índice de um valor no slice
func Index(slice []string, valor string) int {
	for i, item := range slice {
		if item == valor {
			return i
		}
	}
	return -1
}

// Função de contains para checar em um slice
func ContainsTabuleiro(slice []string, valor string) bool {
	for _, simbolo := range slice {
		if simbolo == valor {
			return true
		}
	}
	return false
}

// Função para verificar se existe um espaço vazio
func Contains(tabuleiro map[string][]string, linha string) bool {
	for _, simbolo := range tabuleiro[linha] {
		if simbolo == " " {
			return true
		}
	}
	return false
}

// Função para contar quantos valores iguais ao jogador existem na linha
func CountValor(tabuleiro map[string][]string, chave string, valor string) int {
	contagem := 0
	linha := tabuleiro[chave]
	for _, item := range linha {
		if item == valor {
			contagem++
		}
	}
	return contagem
}