build: clean
	env GOROOT=$(HOME)/go-projects/src/wasm-test/go/ GOARCH=wasm GOOS=js go1.11beta1 build -o example.wasm main.go

clean:
	rm -f example.wasm

godoc:
	env GOROOT=$(HOME)/go-projects/src/wasm-test/go/ godoc -http=:6060

server-main:
	go build -o server-main server/main.go

run-server: server-main
	./server-main
