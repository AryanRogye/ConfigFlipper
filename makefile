OS_NAME := $(shell uname -s | tr A-Z a-z)
BIN := ConfigFlipper

build:
	go build cmd/configflipper/main.go
	mv main ${BIN}
run_debug:
	dlv debug cmd/configflipper/main.go \
		--headless \
		--listen=127.0.0.1:2345 \
		--api-version=2 \
		--accept-multiclient
attach_debugger:
	dlv connect 127.0.0.1:2345

run:
	go run cmd/configflipper/main.go

check_race:
	go run -race cmd/configflipper/main.go

install: build
ifeq ($(OS_NAME),darwin)
	@echo "Installing to /usr/local/bin/$(BIN)"
	sudo mv $(BIN) /usr/local/bin/$(BIN)
else
	@echo "install only supported on macOS right now"
endif
