FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app



FROM alpine:3.20 AS production

COPY --from=builder /go/bin/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]
