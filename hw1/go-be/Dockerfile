FROM golang:1.16 as build

WORKDIR /go/src/app
COPY . .
RUN go test ./test
RUN go build -o /go-be

FROM alpine:3.11.3
COPY --from=build go-be $GOPATH/bin/go-be
#needed because alpine only has musl
RUN apk add libc6-compat 

CMD ["go-be"]