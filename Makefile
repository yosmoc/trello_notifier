BIN := trello_notifier

.PHONY: all
all: clean build

.PHONY: clean
clean:
	rm -rf $(BIN) goxz

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: install
install:
	go install .

.PHONY: crossplatform
crossplatform:
	goxz -os=linux,darwin -n $(BIN) .