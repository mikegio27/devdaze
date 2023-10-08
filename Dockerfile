FROM go:latest

COPY server .

ENTRYPOINT [ "server" ]