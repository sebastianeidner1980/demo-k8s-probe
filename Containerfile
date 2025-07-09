# Stage 1: Build statically linked Go binary
FROM golang:1.22 as builder

WORKDIR /app
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o healthcheck .

# Stage 2: Use distroless base image
FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=builder /app/healthcheck .

USER nonroot
EXPOSE 8080
ENTRYPOINT ["/app/healthcheck"]
