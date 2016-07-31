build:
	go-bindata -pkg handler -o internal/handler/bindata.go assets/...
	go test ./...
	go build ./...

run:
	go-bindata -pkg handler -o internal/handler/bindata.go assets/...
	go test ./...
	go build ./...
	go run main.go