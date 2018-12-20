build:
	@rm -f bin/kx
	@go build -o bin/kx

install:
	@go install

release:
	@git commit -am "Incrementing $(part) version from $(shell waypoint latest) ->"
