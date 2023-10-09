FROM golang:latest


COPY server .

ENTRYPOINT [ "./server" ]