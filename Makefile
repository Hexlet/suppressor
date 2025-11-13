install:
	go mod download

test:
	go test -count=1 .

tidy:
	go mod tidy

update-deps:
	go get -u all && go mod tidy
