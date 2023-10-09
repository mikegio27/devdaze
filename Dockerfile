FROM golang:latest

COPY src/server .

ENTRYPOINT [ "server" ]