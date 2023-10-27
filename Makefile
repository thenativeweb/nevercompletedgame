#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

# Wir definieren für das help-Target eine Ausgabe mit den vorher definierten Farben.
# An den Targets definieren wir mit ##@ eine Gruppierung und nach dem ersten Leerzeichen den Hilfetext.
HELP_FUN = \
        %help; \
        while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
        print "usage: make [target]\n\n"; \
        for (sort keys %help) { \
        print "${WHITE}$$_:${RESET}\n"; \
        for (@{$$help{$$_}}) { \
        $$sep = " " x (32 - length $$_->[0]); \
        print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
        }; \
        print "\n"; }

.DEFAULT_GOAL := help # Normalerweise ist das aufgerufene Target im Makefile das zuerst definierte. Hiermit geben wir es
# vor und es ist egal an welcher Stelle im Code es steht.

help: ##@project Show this help
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)
.PHONY: help
#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

# Wir definieren für das help-Target eine Ausgabe mit den vorher definierten Farben.
# An den Targets definieren wir mit ##@ eine Gruppierung und nach dem ersten Leerzeichen den Hilfetext.
HELP_FUN = \
        %help; \
        while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
        print "usage: make [target]\n\n"; \
        for (sort keys %help) { \
        print "${WHITE}$$_:${RESET}\n"; \
        for (@{$$help{$$_}}) { \
        $$sep = " " x (32 - length $$_->[0]); \
        print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
        }; \
        print "\n"; }

# Normalerweise ist das aufgerufene Target im Makefile das zuerst definierte. Hiermit geben wir es explizit
# vor und es ist egal an welcher Stelle es im Code steht. Die ganz Faulen (wie ich) machen sich noch einen Alias von "m"
# auf "make". ;)
.DEFAULT_GOAL := help

# Wenn der Developer das Projekt nicht kennt und einfach nur "make" eingibt bekommt er nun eine hilfreiche Ausgabe, die
# ihn anleitet.
help: ##@project Show this help
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)
.PHONY: help

# In all unseren Projekten gibt es immer das Target "setup", welches das Projekt in einem definierten Zustand aufsetzt.
# Der Workflow ist damit immer gleich. Du checkst das Projekt per "git clone" aus, machst danach ein "make setup"
# und kannst Dich darauf verlassen, dass das Projekt danach funktioniert. KISS!
# Was genau zum Setup notwendig ist interessiert Dich nicht mehr. It just works.
# Es hat sich auch bewährt die Targets mit gleichem Zweck in allen Projekten immer gleich zu benennen und so einen
# Standard zu etablieren.
setup: build ##@project Set up project

analyze: ##@development Analyze code
setup: build ##@project Set up project

analyze: ##@development Analyze code
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

build: build-docker ##@project Build project
	@go build -o ./build/ncg .

build-docker: ##@project Build docker image used by project
	@docker build -t thenativeweb/ncg .

coverage: test ##@development Calculate code coverage html pages and save to ./coverage/coverage.html file
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/coverage.out ./...
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

test: ##@development Execute go tests including coverage output in percentage
	@go test -cover ./...

tests: test coverage ##@development Execute go tests, calculate code coverage html pages and analyze code

.PHONY: analyze \
	build \
	build-docker \
	coverage \
	setup \
	test
