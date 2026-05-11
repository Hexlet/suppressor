install:
	go mod download

test:
	go test -count=1 .

release:
	git tag $(shell svu patch)
	git push origin $(shell svu current)

tidy:
	go mod tidy

update-deps:
	go get -u all && go mod tidy
