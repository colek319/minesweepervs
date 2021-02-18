# minesweepervs

A multiplayer minesweeper implemented in WebAssembly

To compile, run from root:

```
GOOS=js GOARCH=wasm go build -o assets/minesweeper.wasm cmd/wasm/main.go
```

Install gin to run the server by running:
`go get -u github.com/gin-gonic/gin`

To run the server:
```
> go run cmd/server/main.go
```

And navigate to `localhost:9090`
