# syntax=docker/dockerfile:1
# Build from the repository root so the local proto module is in the context:
#   docker build -f payment/Dockerfile -t microservices/payment:dev .

# ---- build stage: compile a static binary ----
FROM golang:1.27.3-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOWORK=off
WORKDIR /src

# Download modules first, so code changes don't invalidate this layer.
COPY proto/go.mod proto/go.sum ./proto/
COPY payment/go.mod payment/go.sum ./payment/
WORKDIR /src/payment
RUN --mount=type=cache,target=/go/pkg/mod go mod download

WORKDIR /src
COPY proto/ ./proto/
COPY payment/ ./payment/
WORKDIR /src/payment
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/payment ./cmd

# ---- runtime stage: just the binary, running as a non-root user ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /out/payment /payment
EXPOSE 8081
ENTRYPOINT ["/payment"]
