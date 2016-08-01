FROM poorny/base

RUN go get github.com/jyotiska/go-webcolors
RUN go get github.com/nfnt/resize

ADD . /go/src/beta.shelfgit.com/mdata/metaimage
WORKDIR /go/src/beta.shelfgit.com/mdata/metaimage

RUN godep get

RUN go install beta.shelfgit.com/mdata/metaimage

CMD ["go", "run", "main.go"]
