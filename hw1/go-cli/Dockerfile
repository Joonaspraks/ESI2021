FROM golang:1.16 as build

WORKDIR /go/src/app
COPY . .

RUN go build -o /go-cli

FROM alpine:3.11.3
COPY --from=build go-cli $GOPATH/bin/go-cli
#needed because alpine only has musl
RUN apk add libc6-compat 

CMD ["/bin/sh"]