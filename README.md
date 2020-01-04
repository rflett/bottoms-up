# Bottoms Up

A gin-gonic API that returns drinking games.

## Development

### Requirements

- Go 1.13+
- Docker

Install Go dependencies:

```bash
go get github.com/gin-gonic/gin
```

Then run with `go run sip.go`.

## Usage

```bash
curl http://localhost:9000/game

{
  "name": "Thumbs",
  "description": "Last person to put their thumb on the table takes a drink",
  "dice_role": 2
}
```
