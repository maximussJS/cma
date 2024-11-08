FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000

CMD ["go", "run", "cmd/server/main.go"]
