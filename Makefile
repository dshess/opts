# This Makefile is vaguely based on the Alex Edwards Makefile.  I think?
# I've kind of lost the thread at this point.  "make help" kind of works.
# I mostly use "make test/live", "make tidy" and "make audit".
#
# Why a Makefile at all?  Because otherwise I just end up with a file full
# of command-lines to copy/paste.
main_package_path = ./

# Files monitored by /live targets.
livemonitor = *.go Makefile

# https://github.com/eradman/entr
# Check your system package manager first, though.
# TODO: See if wgo is a better option for entr.
.PHONY: live
live:
	ls -1d ${livemonitor} | entr -r -s "make $(LIVETARGET)"

# ============================================================================ #
# HELPERS
# ============================================================================ #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)"


# ============================================================================ #
# QUALITY CONTROL
# ============================================================================ #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1020,-ST1021,-ST1022,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/live: make test every time a file changes
.PHONY: test/live
test/live:
	make LIVETARGET=test live

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out


# ============================================================================ #
# DEVELOPMENT
# ============================================================================ #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...


# ============================================================================ #
# OPERATIONS
# ============================================================================ #

## push: push changes to git origin
.PHONY: push
push: confirm audit no-dirty
	git push

## push/github: push changes to github
.PHONY: push/github
push/github: confirm audit no-dirty
	git push/github

## push/all: push changes to all git remotes
.PHONY: push/all
push/all: confirm audit no-dirty
	for remote in `git remote`; do /bin/echo -n $$remote: ""; git push $$remote; done
