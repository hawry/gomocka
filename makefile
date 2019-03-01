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

static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(OUT)s $(LDFLAGS)