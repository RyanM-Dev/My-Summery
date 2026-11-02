# syntax=docker/dockerfile:1
# Build from the repository root so the local proto module is in the context:
#   docker build -f order/Dockerfile -t microservices/order:dev .

# ---- build stage: compile a static binary ----
FROM golang:1.27.3-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOWORK=off
WORKDIR /src

# Download modules first, so code changes don't invalidate this layer.
COPY proto/go.mod proto/go.sum ./proto/
COPY order/go.mod order/go.sum ./order/
WORKDIR /src/order
RUN --mount=type=cache,target=/go/pkg/mod go mod download

WORKDIR /src
COPY proto/ ./proto/
COPY order/ ./order/
WORKDIR /src/order
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /out/order ./cmd

# ---- runtime stage: just the binary, running as a non-root user ----
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /out/order /order
EXPOSE 8080
ENTRYPOINT ["/order"]
