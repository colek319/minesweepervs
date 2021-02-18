FROM golang:1.15 
WORKDIR /minesweepervs

EXPOSE 9090

COPY go.mod go.sum /minesweepervs/
RUN go mod download

COPY . .

RUN GOOS=js GOARCH=wasm go build -o assets/minesweeper.wasm cmd/wasm/main.go 

CMD go run cmd/server/main.go
