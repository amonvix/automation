## 🎒 README (descontraído – pra você treinar e manter ritmo)

```md
# Automation Gym (Go)

This repo is my personal training ground for automation in Go.  
If I can automate it, I will. If I can’t, I’ll learn how.

## Why this exists

Because repeating the same boring stuff manually is how humans create bugs.
This is my way of turning “I’ll remember this command” into “I’ll never type this again”.

## Tools I’m training with

### standard
The “hello world” of automation.

Flags:
- `--path`: where the chaos lives
- `--dry-run`: look, don’t touch

Examples:
```bash
go run ./cmd/standard --path ./src --dry-run
ENV=dev go run ./cmd/standard --path ./src
