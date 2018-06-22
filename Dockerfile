FROM golang
ADD . /go/src/github.com/mikenavarro/scratch-go
RUN go install github.com/mikenavarro/scratch-go
ENTRYPOINT ["/go/bin/scratch-go"]