# T2-FPPD
Implementação de Jogo Multiplayer Usando RPC em Go


jogo-multiplayer-go/
│
├── server/
│   └── main.go          # Código do servidor RPC
│
├── client/
│   └── main.go          # Código do cliente (interface e lógica do jogo)
│
├── models/
│   ├── player.go        # Estrutura Player
│   └── gamestate.go     # Estrutura GameState
│
├── rpc/
│   ├── server.go        # Métodos RPC do servidor
│   └── client.go        # Métodos RPC do cliente
│
├── go.mod
└── README.md