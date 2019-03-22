######## Build image #######
FROM golang:1.12.1 as builder
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -u github.com/nlopes/slack
LABEL version="0.1"
ADD . /go/src/ask8ball
WORKDIR /go/src/ask8ball
RUN dep ensure
RUN CGO_ENABLED=1 go build -a -v
RUN CGO_ENABLED=1 GOOS=linux go install -v
ENTRYPOINT ["/go/bin/ask8ball"]


FROM alpine:3.9.2
RUN apk --no-cache add ca-certificates
WORKDIR /app
RUN cd /app
RUN apk add --no-cache libc6-compat
COPY --from=builder /go/bin/ask8ball /go/bin/ask8ball
ENTRYPOINT ["/go/bin/ask8ball"]