######## Build image #######
FROM golang:1.12.1 as builder
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/nlopes/slack
LABEL version="0.1"
ADD . /go/src/ask8ball
WORKDIR /go/src/ask8ball
RUN dep ensure
RUN go build
RUN go install
ENTRYPOINT ["/go/bin/ask8ball"]

#FROM scratch
#COPY --from=builder /go/bin/ask8ball /go/bin/ask8ball
#ENTRYPOINT ["/go/bin/ask8ball"]