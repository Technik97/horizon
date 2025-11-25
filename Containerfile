# --- Build stage ---
FROM docker.io/library/golang:1.25-alpine AS builder

WORKDIR /src
# Copy the Go module files first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build Horizon statically
RUN go build -o /horizon main.go

# --- Final stage ---
FROM scratch
# Copy the binary from the builder stage
COPY --from=builder /horizon /horizon

# Set binary as entrypoint
ENTRYPOINT ["/horizon"]
CMD []
