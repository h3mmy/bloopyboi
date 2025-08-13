FROM golang:1.25-alpine@sha256:72a77777b4d52d80820a9b5828c06cca1cd3c24721b6ca25b0269b5f6fb00daa as build

ARG TARGETPLATFORM
ENV TARGETPLATFORM=${TARGETPLATFORM:-linux/amd64}

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN export GOOS=$(echo ${TARGETPLATFORM} | cut -d / -f1) \
    && \
    export GOARCH=$(echo ${TARGETPLATFORM} | cut -d / -f2) \
    && \
    GOARM=$(echo ${TARGETPLATFORM} | cut -d / -f3); export GOARM=${GOARM:1}
RUN go mod download
# These are done in a different part of the pipeline. 
# RUN go vet -v
# RUN go test -v ./...
RUN go build -ldflags="-w -s" .
RUN echo $(ls .)

FROM gcr.io/distroless/static@sha256:2e114d20aa6371fd271f854aa3d6b2b7d2e70e797bb3ea44fb677afec60db22c

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/bloopyboi"]
