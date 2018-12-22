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

ifneq ($(DIRTY), )
	$(error git working tree is dirty; commit or stash changes and try again)
endif
	./do_release_commit.sh $(part)
