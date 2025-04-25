package jogadaBot

import (
	"jogo_da_velha/utils"
	"math/rand"
	"time"
)

func JogadaNivelFacil(numeroUtilizado []int) int {

	fonte := rand.NewSource(time.Now().UnixNano())
	r := rand.New(fonte)

	jogadaBot := r.Intn(9)

	for utils.JaUtilizado(jogadaBot, numeroUtilizado) {
		jogadaBot = r.Intn(9)
	}

	return jogadaBot
}