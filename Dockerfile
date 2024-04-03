FROM golang:1.22-alpine@sha256:cdc86d9f363e8786845bea2040312b4efa321b828acdeb26f393faa864d887b0 as build

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

FROM gcr.io/distroless/static@sha256:046b92c933032a8ca99a66f4c79a68ac029d9a4ababd1a806a82140b3b899fd3

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/bloopyboi"]
