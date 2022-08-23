DIRTY := $(shell git diff --stat)

build:
	@rm -f .bin/kx
	@go build -o .bin/kx

install:
	@go install

## Create a virtualenv and install package for managing versions
installbump:
ifeq ($(wildcard ./.venv/.*),)
	@python3 -m venv ./.venv
	@./.venv/bin/pip install -U pip git+https://github.com/kylie-a/bumpversion.git
endif

release: installbump
ifndef part
	$(error 'part' is undefined. Use 'major', 'minor' or 'patch' as the part argument)
endif

ifneq ($(DIRTY), )
	$(error git working tree is dirty; commit or stash changes and try again)
endif
	./do_release_commit.sh $(part)
