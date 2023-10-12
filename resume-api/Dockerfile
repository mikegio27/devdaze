FROM golang:latest


COPY server .
COPY utils/resume/resume.json .
EXPOSE 8080
ENTRYPOINT [ "./server" ]