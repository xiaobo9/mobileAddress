# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags " -w" -o /mobileAddress

##
## Deploy
##
FROM scratch

WORKDIR /

COPY data /data
COPY static /static
COPY --from=build /mobileAddress /mobileAddress

EXPOSE 8080

ENTRYPOINT ["/mobileAddress"]
