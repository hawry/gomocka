OUT=gock
GITDESC=`git describe --always`
GITCOUNT=`git rev-list --count --first-parent HEAD`
LDFLAGS=-ldflags "-X main.buildVersion=$(GITDESC)-$(GITCOUNT)"
RUNFLAGS=--verbose

.PHONY: all
.SILENT:

all: build run

build:
	go build -o $(OUT) $(LDFLAGS)

run:
	./$(OUT) $(RUNFLAGS)

clean:
	rm -rf ./$(OUT)