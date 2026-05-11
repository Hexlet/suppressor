install:
	go mod download

test:
	go test -count=1 .

release:
	@NEXT=$$(svu patch) && git tag $$NEXT && git push origin $$NEXT

tidy:
	go mod tidy

update-deps:
	go get -u all && go mod tidy
