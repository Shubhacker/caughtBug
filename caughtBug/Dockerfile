FROM golang:1.17.11-alpine as build-env

RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories
RUN echo http://dl-cdn.alpinelinux.org/alpine/edge/testing >> /etc/apk/repositories
RUN apk update

RUN apk --no-cache add gcc g++ make git

## Update CA Certificates
RUN apk add --no-cache git ca-certificates && update-ca-certificates

# Download all timezone data
RUN apk add --no-cache tzdata
RUN apk add tzdata

# Add Maintainer Info
LABEL maintainer="Caught Bug"

# Set the Current Working Directory inside the container
RUN mkdir /app
WORKDIR /app

COPY . .

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/caughtbug

FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/bin/caughtbug /go/bin/caughtbug
COPY --from=build-env /app/ /
COPY --from=build-env /usr/share/zoneinfo/ /usr/share/zoneinfo/

ENTRYPOINT ["/go/bin/caughtbug", "docker"]
