FROM golang:1.20-alpine
ENV GO111MODULE on

RUN apk update && \
    apk --no-cache add git && \
    apk add bash
RUN go install github.com/cosmtrek/air@v1.29.0

RUN mkdir /go/src/app
WORKDIR /go/src/app

CMD ["air", "-c", ".air.toml"]