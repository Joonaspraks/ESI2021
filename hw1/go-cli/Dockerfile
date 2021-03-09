FROM golang:1.16

WORKDIR /go/src/app
COPY . .

#RUN go install
RUN go build -o $GOPATH/bin/go-cli
#CMD ["app"]
#CMD ["go", "build", "-o", "$GOPATH/bin/go-cli.exe"]
CMD ["/bin/bash"]