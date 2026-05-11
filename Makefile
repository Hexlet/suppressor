install:
	go mod download

test:
	go test -count=1 .

release:
	git tag $(shell svu next)
	git push origin $(shell git describe --tags --abbrev=0)

tidy:
	go mod tidy

update-deps:
	go get -u all && go mod tidy
