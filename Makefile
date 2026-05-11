install:
	go mod download

test:
	go test -count=1 .

release:
	@NEXT=$(shell svu next); \
	CURRENT=$(shell svu current); \
	if [ "$$NEXT" = "$$CURRENT" ]; then \
		echo "Nothing to release: no new commits since $$CURRENT"; \
		exit 1; \
	fi; \
	git tag $$NEXT && git push origin $$NEXT

tidy:
	go mod tidy

update-deps:
	go get -u all && go mod tidy
