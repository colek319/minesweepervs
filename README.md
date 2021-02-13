# minesweepervs

A multiplayer minesweeper implemented in WebAssembly

To compile, run from root:

```
GOOS=js GOARCH=wasm go build -o assets/minesweeper.wasm cmd/wasm/main.go
```

To run the server:
```
> cd cmd/server
> go run main.go
```

And navigate to `localhost:9090`