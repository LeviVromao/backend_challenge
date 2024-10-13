FROM golang:1.19-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN ls ./cmd

RUN go build -o /app/bin/server ./cmd/api/main.go

EXPOSE 8080

CMD ["/app/bin/server"]
