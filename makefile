OUT=gomocka
PACKAGES = $(shell find ./ -type d -not -path '*/\.*')
GITDESC=`git describe --always --tags`
GITCOUNT=`git rev-list --count --first-parent HEAD`
LDFLAGS=-ldflags "-X main.buildVersion=$(GITDESC)-$(GITCOUNT)"
RUNFLAGS=--verbose

.PHONY:
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

release: static
	mkdir -p build; \
	tar -cvzf ./build/gomocka-linux-amd64.tar.gz $(OUT)

static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(OUT) $(LDFLAGS)

report:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	BROWSER=firefox go tool cover -html=coverage-all.out