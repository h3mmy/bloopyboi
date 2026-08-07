FROM golang:1.25-alpine@sha256:8e02eb337d9e0ea459e041f1ee5eece41cbb61f1d83e7d883a3e2fb4862063fa AS build

ARG TARGETPLATFORM
ENV TARGETPLATFORM=${TARGETPLATFORM:-linux/amd64}

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build go mod download

COPY . .

RUN export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) \
    && \
    export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) \
    && \
    GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1}
# These are done in a different part of the pipeline.
# RUN go vet -v
# RUN go test -v ./...
RUN --mount=type=cache,target=/root/.cache/go-build go build -ldflags="-w -s" .
RUN echo $(ls .)

FROM gcr.io/distroless/static@sha256:d90359c7a3ad67b3c11ca44fd5f3f5208cbef546f2e692b0dc3410a869de46bf

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000
EXPOSE 8080

ENTRYPOINT ["/bloopyboi"]
