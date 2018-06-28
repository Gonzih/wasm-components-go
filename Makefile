build: clean go
	env GOROOT=$(HOME)/go-projects/src/wasm-test/go/ GOARCH=wasm GOOS=js go1.11beta1 build -o example.wasm *.go

clean:
	rm -f example.wasm

godoc:
	env GOROOT=$(HOME)/go-projects/src/wasm-test/go/ godoc -http=:6060

server-main:
	go build -o server-main server/main.go

run-server: server-main
	./server-main

setup:
	go get golang.org/x/build/version/go1.11beta1
	go1.11beta1 download

go:
	wget https://dl.google.com/go/go1.11beta1.src.tar.gz -O /tmp/go1.11beta1.src.tar.gz
	tar xvzf /tmp/go1.11beta1.src.tar.gz
