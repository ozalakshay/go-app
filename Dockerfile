# Build stage
FROM golang:1.22.5 AS builder

WORKDIR /app

# Copy only what's needed first
COPY go.mod ./
RUN go mod download

# Now copy everything
COPY . .

# Build the Go app as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage: small secure image
FROM gcr.io/distroless/static

WORKDIR /

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

EXPOSE 8080

ENTRYPOINT ["/main"]
