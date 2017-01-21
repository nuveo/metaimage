FROM golang:1.7

RUN git clone https://github.com/nuveo/metaimage.git /go/src/github.com/nuveo/metaimage
WORKDIR /go/src/github.com/nuveo/metaimage

RUN go get -u github.com/kardianos/govendor
RUN govendor fetch +e

CMD ["go", "run", "main.go"]
