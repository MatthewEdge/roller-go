build:
	go build -o bin/roll main.go

test:
	go test ./...

cover:
	go test -coverprofile=cover.tmp ./...
	go tool cover -html=cover.tmp -o cover.html
	rm -f cover.tmp
