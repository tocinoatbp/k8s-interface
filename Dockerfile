FROM golang:1.13.8-alpine3.11
LABEL maintainer="tocinoatbp"
RUN apk update --no-cache && \
    apk add curl
COPY . .
RUN go build .
CMD [ "./k8s-interface" ]