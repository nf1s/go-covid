test:
	go test
	

test-coverage:
	go test -coverprofile=c.out
    go tool cover -html=c.out -o coverage.html

deploy-docs:
	mkdocs gh-deploy