run:
	go run main.go
	
test:
	 go test ./... --cover 

format:
	 go fmt ./...