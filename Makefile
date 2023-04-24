generate:
	@echo "Generating..."
	sh ./gen_script.sh
	@echo "Done"

lintci-deps:
	@echo "Installing linters..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./build/bin latest
	@echo "Done"

lint:
	@echo "Linting..."
	golangci-lint run --fix
	@echo "Done"

test: 
	@echo "Testing..."
	go test -v ./...
	@echo "Done"