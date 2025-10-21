package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

var rpcClient *rpc.Client
var localPlayer PlayerState
var sequenceNumber int = 0
var otherPlayers map[int]PlayerState

// Conecta ao servidor RPC e registra o jogador
func conectarServidor(nome string) {
	client, err := rpc.Dial("tcp", "localhost:8932")
	if err != nil {
		log.Fatal("Erro ao conectar ao servidor:", err)
	}
	rpcClient = client
	fmt.Println("✅ Conectado ao servidor RPC!")

	var player PlayerState
	err = rpcClient.Call("GameService.RegisterPlayer", &nome, &player)
	if err != nil {
		log.Fatal("Erro ao registrar jogador:", err)
	}
	localPlayer = player
	fmt.Printf("Jogador registrado: ID=%d, Nome=%s\n", localPlayer.ID, localPlayer.Name)
}

// Envia o estado atual do jogador ao servidor
func enviarEstado(posX, posY int) {
	sequenceNumber++
	localPlayer.PosX = posX
	localPlayer.PosY = posY
	localPlayer.Sequence = sequenceNumber

	var ok bool
	for {
		err := rpcClient.Call("GameService.UpdatePlayer", &localPlayer, &ok)
		if err == nil {
			break
		}
		fmt.Println("⚠️ Falha ao enviar estado, tentando novamente...")
		time.Sleep(500 * time.Millisecond)
	}
}

// Goroutine: sincroniza o estado do jogo (busca outros jogadores)
func sincronizarEstado() {
	for {
		var gs map[int]PlayerState
		err := rpcClient.Call("GameService.GetGameState", &struct{}{}, &gs)
		if err == nil {
			otherPlayers = gs
		}
		time.Sleep(300 * time.Millisecond)
	}
}
