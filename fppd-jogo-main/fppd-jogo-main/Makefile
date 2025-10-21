.PHONY: all build

all: build
	go build client.go
	go build server.go

go.mod:
	go mod init jogo
	go get -u github.com/nsf/termbox-go

build: go.mod
	go build -o jogo
	
clean:
	rm -f jogo
	rm -f client server *.exe

distclean: clean
	rm -f go.mod go.sum
	
