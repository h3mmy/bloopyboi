FROM golang:1.25-alpine@sha256:f6751d823c26342f9506c03797d2527668d095b0a15f1862cddb4d927a7a4ced AS build

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

FROM gcr.io/distroless/static@sha256:eca24e67792afe660769e84c85332d9939772e7f76071597d179af96ac4e9e4f

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000
EXPOSE 8080

ENTRYPOINT ["/bloopyboi"]
