// main.go - Loop principal do jogo
package main

import "os"

func main() {
	interfaceIniciar()
	defer interfaceFinalizar()

	conectarServidor("Luana") // nome do jogador
	go sincronizarEstado()    // atualiza outros jogadores periodicamente

	mapaFile := "mapa.txt"
	if len(os.Args) > 1 {
		mapaFile = os.Args[1]
	}

	jogo := jogoNovo()
	if err := jogoCarregarMapa(mapaFile, &jogo); err != nil {
		panic(err)
	}

	interfaceDesenharJogo(&jogo)

	for {
		evento := interfaceLerEventoTeclado()
		if continuar := personagemExecutarAcao(evento, &jogo); !continuar {
			break
		}
		enviarEstado(jogo.PosX, jogo.PosY)
		interfaceDesenharJogo(&jogo)
	}
}
