# Build Env
FROM golang:1.12 AS build-env

ENV GO111MODULE=on

ADD . /go/src/github.com/blockchain-abstraction-middleware/project-template

WORKDIR /go/src/github.com/blockchain-abstraction-middleware/project-template

RUN go build -i -o app .

# Application Image
FROM gcr.io/distroless/base:latest

COPY --from=build-env /go/src/github.com/blockchain-abstraction-middleware/project-template/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]