package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// GameService gerencia o estado global dos jogadores
type GameService struct {
	mu        sync.Mutex
	players   map[int]PlayerState // estado atual dos jogadores
	processed map[int]int         // controle de comandos processados (exactly-once)
	nextID    int
}

// Registra um novo jogador e retorna seu ID
func (g *GameService) RegisterPlayer(name *string, reply *PlayerState) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.nextID++
	player := PlayerState{
		ID:   g.nextID,
		Name: *name,
		PosX: 0,
		PosY: 0,
	}

	g.players[player.ID] = player
	*reply = player

	fmt.Printf("[Server] Jogador registrado: %+v\n", player)
	return nil
}

// Atualiza a posição e o estado de um jogador (chamado pelo cliente)
func (g *GameService) UpdatePlayer(state *PlayerState, reply *bool) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Garantia de execução única (exactly-once)
	if seq, ok := g.processed[state.ID]; ok && state.Sequence <= seq {
		*reply = true
		return nil
	}

	g.processed[state.ID] = state.Sequence
	g.players[state.ID] = *state

	fmt.Printf("[Server] Atualização recebida: %+v\n", *state)
	*reply = true
	return nil
}

// Retorna o estado completo do jogo (todos os jogadores)
func (g *GameService) GetGameState(_ *struct{}, reply *map[int]PlayerState) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	gameCopy := make(map[int]PlayerState)
	for id, p := range g.players {
		gameCopy[id] = p
	}

	*reply = gameCopy
	fmt.Printf("[Server] Estado enviado (%d jogadores)\n", len(gameCopy))
	return nil
}

func mainServer() {
	game := &GameService{
		players:   make(map[int]PlayerState),
		processed: make(map[int]int),
		nextID:    0,
	}

	rpc.Register(game)

	listener, err := net.Listen("tcp", ":8932")
	if err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
	fmt.Println("Servidor RPC do jogo iniciado na porta :8932")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
