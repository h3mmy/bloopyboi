FROM golang:1.22-alpine@sha256:eba0670e527735bed965a32217e5db1f895c544619d726cccb62b1be6d11f287 as build

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

FROM gcr.io/distroless/static@sha256:9be3fcc6abeaf985b5ecce59451acbcbb15e7be39472320c538d0d55a0834edc

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/bloopyboi"]
