DIRTY := $(shell git diff --stat)

build:
	@rm -f bin/kx
	@go build -o bin/kx

install:
	@go install

release:
ifndef part
	$(error 'part' is undefined. Use 'major', 'minor' or 'patch' as the part argument)
endif

ifneq ($(DIRTY), "")
	$(error git working tree is dirty; commit or stash changes and try again)
endif
		@git commit --allow-empty -m "Incrementing $(part) version from $(shell waypoint latest kx) -> $(shell waypoint bump kx --$(part))"
		@git push origin master
		@git tag $(shell waypoint latest kx)
		@git push --tags
