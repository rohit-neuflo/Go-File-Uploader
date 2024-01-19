PORT := 3001
OS := $(shell uname)

build:
	@go build -o bin/uploader

run: build
	@if [ "$(OS)" = "Darwin" ]; then \
		make kill_old_process; \
	fi
	@./bin/uploader

kill_old_process:
	@echo "Checking for existing process on port $(PORT) and killing it if found..."
	@PID=$$(lsof -ti :$(PORT)); \
	if [ -n "$$PID" ]; then \
		echo "Killing existing process on port $(PORT) with PID $$PID"; \
		kill -9 $$PID; \
	fi

test:
	@go test -v ./...
