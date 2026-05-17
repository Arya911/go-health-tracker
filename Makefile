run:
	go run cmd/healthcmd/main.go
build:
	go build cmd/healthcmd/main.go
	./main
clean:
	rm -rf main
