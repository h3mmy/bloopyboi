FROM golang:1.22-alpine@sha256:0f38b77b150dc9bd784d90cc7942a4cfd712242711fdf490f6ae85d5040e2c82 as build

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

FROM gcr.io/distroless/static@sha256:072d78bc452a2998929a9579464e55067db4bf6d2c5f9cde582e33c10a415bd1

COPY --from=build /build/bloopyboi /

WORKDIR /

EXPOSE 3000

ENTRYPOINT ["/bloopyboi"]
