FROM golang:1.20.1-alpine3.17


RUN apk add build-base ; \
    go install github.com/go-delve/delve/cmd/dlv@latest 

COPY ./docker/dev/entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
