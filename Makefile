run: ./cmd/main.go
	go run ./cmd/main.go

build: ./cmd/main.go
	go build ./cmd/main.go -o ./build/the-chemical-apocalypse
