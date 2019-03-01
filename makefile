OUT=gock
LDLFLAGS=
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