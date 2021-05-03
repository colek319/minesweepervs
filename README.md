# minesweepervs


## Local Build (non-docker)
A multiplayer minesweeper implemented in WebAssembly

To compile, run from root:

```
GOOS=js GOARCH=wasm go build -o assets/minesweeper.wasm cmd/wasm/main.go
```

Install gin-gonic and ws library to run the server by running:
`go get -u github.com/gin-gonic/gin && go get nhooyr.io/websocket`

To run the server:
```
> go run cmd/server/main.go
```


## Docker Build 
Alternatively, use docker-compose. Run:

```
docker-compose up
```

This will run the docker run configuration for every service defined in docker-compose.yml. 

Navigate to `localhost:9090`
