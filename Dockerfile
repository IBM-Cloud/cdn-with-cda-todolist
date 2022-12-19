FROM golang:buster
WORKDIR /go/src/app
ADD . /go/src/app
ADD ./views /views
ADD ./static /static
RUN go get -d -v ./...
RUN go build -ldflags="-s -w" -o /app
CMD ["/app"]
EXPOSE 8080