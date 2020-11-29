FROM golang:1.13.8-alpine3.11
RUN apk update --no-cache && \
    apk add curl
COPY . .
# CMD ["/bin/sh","-c","sleep 9999"]
CMD [ "/usr/local/go/bin/go","run", "main.go" ]