OUT=gomocka
GITDESC=`git describe --always --tags`
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

test:
	go test ./... -v

static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(OUT) $(LDFLAGS)
