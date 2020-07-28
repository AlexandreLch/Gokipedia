# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.14.1-alpine3.11 as builder

ENV GO111MODULE=on

# The latest alpine image don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Equipe 01 - Hetic"

WORKDIR /go/src/gokipedia
COPY go.sum go.mod /go/src/gokipedia/

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . /go/src/gokipedia/

# Install and run packr
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN packr2

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gokipedia .

FROM alpine

COPY --from=builder /go/src/gokipedia/bin/gokipedia /

EXPOSE 8080

ENTRYPOINT ["/gokipedia"]