FROM golang:1.24.2 AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]
