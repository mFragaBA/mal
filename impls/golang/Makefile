export GOPATH := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

#####################

SRCS = step0_repl.go
BINS = $(SRCS:%.go=%)

#####################

all: $(BINS)

step0_repl: src/step0_repl.go
	go build src/step0_repl.go

clean:
	rm -f $(BINS)

