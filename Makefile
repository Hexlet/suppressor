install:
	go mod download

test:
	go test -count=1 .
