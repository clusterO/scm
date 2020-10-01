FROM golang

RUN mkdir -p /go/src/scm

ADD . /go/src/scm

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run scm/osm/cmd  -watch scm/osm
