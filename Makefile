build:
	@rm -f bin/kx
	@go build -o bin/kx

install:
	@go install

release:
	@git commit -am "Incrementing $(part) version from $(shell waypoint latest kx) -> $(shell waypoint bump kx --$(part))"
	@git push origin master
	@git tag $(shell waypoint latest kx)
	@git push --tags
