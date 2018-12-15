build:
	@rm -f bin/kx
	@go build -o bin/kx

install:
	@go install
