run:
	go run src/main.go
	
test:
	 go test ./... --cover 

format:
	 go fmt ./...