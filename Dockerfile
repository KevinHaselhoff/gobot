FROM golang:1.10-alpine3.7

WORKDIR $GOPATH/src/gobot
COPY . .

RUN apk update && apk add curl git && \
    curl https://glide.sh/get | sh && \
    glide up && glide install

RUN go install ./...

CMD ["gobot"]