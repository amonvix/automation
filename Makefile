.PHONY: standard preflight incident async all help

help:
	@echo "Available targets:"
	@echo "  make standard   -> run standard CLI"
	@echo "  make preflight  -> run preflight checks"
	@echo "  make incident   -> run incident helper"
	@echo "  make async      -> run async runner"
	@echo "  make all        -> run all tools (demo)"

standard:
	go run ./cmd/standard

preflight:
	go run ./cmd/preflight

incident:
	go run ./cmd/incident

async:
	go run ./cmd/async_runner

all: preflight standard incident async
