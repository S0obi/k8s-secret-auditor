FROM golang:1.15-alpine AS build
WORKDIR /src
COPY . .
RUN go get -v -t -d ./... 
RUN go build -o bin/ -v ./...

ENTRYPOINT [ "/src/bin/k8s-secret-auditor" ]