FROM golang:1.21-alpine as build

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
# RUN go vet -v
# RUN go test -v
RUN go build -x -ldflags="-w -s" .
RUN echo $(ls .)

FROM gcr.io/distroless/static

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/bloopyboi"]
