which = $(shell which $1 2> /dev/null || echo $1)

GO_PATH := $(call which,go)

TOOLS=$(wildcard tools/*)
.PHONY: ${TOOLS}

${TOOLS}:
	@echo "running $(subst tools/,,$@)"
	@$(GO_PATH) run ./$@

data/dictionary.txt:
	@make tools/scrape > data/dictionary.txt
.PHONY: data/dictionary.txt
