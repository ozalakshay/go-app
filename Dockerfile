# Build stage
FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage: Distroless
FROM gcr.io/distroless/static

WORKDIR /

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

EXPOSE 8080

ENTRYPOINT ["/main"]
