test:
	cd covid && go test -v
	

test-coverage:
	cd covid && go test -v -coverprofile=c.out
    go tool cover -html=c.out -o coverage.html

deploy-docs:
	mkdocs gh-deploy