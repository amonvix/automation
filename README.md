# Automation Toolbox (Go)

A collection of small, focused CLI tools written in Go to automate common development and incident-response tasks.

This repository is designed as a personal automation toolkit, following production-oriented practices:
- clear CLI interfaces (flags)
- preflight checks
- dry-run support
- structured logging
- synchronous and asynchronous execution patterns

## Tools

### standard
Baseline CLI template with flags and environment awareness.

**Flags**
- `--path` (string): Path to scan or operate on.
- `--dry-run` (bool): Execute without making changes.

**Example**
```bash
go run ./cmd/standard --path ./src --dry-run
ENV=dev go run ./cmd/standard --path ./src
```

For learning notes and experimental patterns, see TRAINING.md