FROM golang:1.10-alpine3.7

WORKDIR $GOPATH/src/gobot
COPY application/ .

RUN apk update && apk add curl git && \
    curl https://glide.sh/get | sh && \
    glide install

RUN go install ./...

CMD ["gobot"]