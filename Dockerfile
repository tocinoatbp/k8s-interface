FROM golang:1.13.8-alpine3.11
RUN apk update --no-cache && \
    apk add curl
COPY . .
RUN go build main.go
CMD [ "./main ]