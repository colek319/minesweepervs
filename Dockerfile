FROM golang:1.15 
WORKDIR /minesweepervs

COPY go.mod go.sum /minesweepervs/
RUN go mod download

COPY cmd pkg assets ./ 

RUN GOOS=js GOARCH=wasm go build -o assets/minesweeper.wasm cmd/wasm/main.go && \
    cd cmd/server && \
    go run main.go

