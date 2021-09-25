BIN := trello_notifier

.PHONY: all
all: clean build

.PHONY: clean
clean:
	rm -rf $(BIN)

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: install
install:
	go install .