FROM golang:latest


COPY server .
COPY utils/resume/resume.json .

ENTRYPOINT [ "./server" ]