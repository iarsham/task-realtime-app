FROM golang:1.22.4 AS builder

WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o /go/bin/main ./cmd/main.go

FROM gcr.io/distroless/static AS prod
WORKDIR /production
COPY --from=builder /go/bin/main .
EXPOSE 8002
ENTRYPOINT ["./main"]