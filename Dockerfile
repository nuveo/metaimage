FROM golang:1.7

ADD . /go/src/shelfgit.com/mdata/metaimage
WORKDIR /go/src/shelfgit.com/mdata/metaimage

RUN go get -u github.com/kardianos/govendor
RUN govendor fetch +e

RUN go install shelfgit.com/mdata/metaimage

CMD ["go", "run", "main.go"]
