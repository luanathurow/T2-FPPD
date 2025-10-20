package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os"
)

type User struct {
    ID    int
    Name  string
    Email string
}

type CreateUserRequest struct {
    Name  string
    Email string
}

type GetUserRequest struct {
    ID int
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Uso: go run client.go <endereço:porta>")
        os.Exit(1)
    }

    serverAddr := os.Args[1] // Ex: "localhost:1234"

    client, err := rpc.Dial("tcp", serverAddr)
    if err != nil {
        log.Fatal("Erro ao conectar:", err)
    }

    // Criar usuário
    createReq := CreateUserRequest{Name: "Maria", Email: "maria@example.com"}
    var newUser User
    err = client.Call("UserService.CreateUser", &createReq, &newUser)
    if err != nil {
        log.Fatal("Erro ao criar usuário:", err)
    }
    fmt.Println("Usuário criado:", newUser)

    // Obter usuário pelo ID
    var fetched User
    getReq := GetUserRequest{ID: newUser.ID}
    err = client.Call("UserService.GetUser", &getReq, &fetched)
    if err != nil {
        log.Fatal("Erro ao buscar usuário:", err)
    }
    fmt.Println("Usuário encontrado:", fetched)

    // Listar todos os usuários
    var allUsers []User
    err = client.Call("UserService.ListUsers", &struct{}{}, &allUsers)
    if err != nil {
        log.Fatal("Erro ao listar usuários:", err)
    }

    fmt.Println("Todos os usuários:")
    for _, u := range allUsers {
        fmt.Printf("ID: %d, Nome: %s, Email: %s\n", u.ID, u.Name, u.Email)
    }
}
