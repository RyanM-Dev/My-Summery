# syntax=docker/dockerfile:1
# Build from the repository root so the local proto module is in the context:
#   docker build -f shipping/Dockerfile -t microservices/shipping:dev .

# ---- build stage: compile a static binary ----
FROM golang:1.27.3-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOWORK=off
WORKDIR /src

# Download modules first, so code changes don't invalidate this layer.
COPY proto/go.mod proto/go.sum ./proto/
COPY shipping/go.mod shipping/go.sum ./shipping/
WORKDIR /src/shipping
RUN --mount=type=cache,target=/go/pkg/mod go mod download

WORKDIR /src
COPY proto/ ./proto/
COPY shipping/ ./shipping/
WORKDIR /src/shipping
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/shipping ./cmd

# ---- runtime stage: just the binary, running as a non-root user ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /out/shipping /shipping
EXPOSE 8082
ENTRYPOINT ["/shipping"]
