FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY ./. .
COPY ./main.go .
COPY ./go.mod .
COPY ./go.sum .
COPY ./handler ./handler
COPY ./model ./model
COPY ./repo ./repo
COPY ./service ./service
COPY ./routing ./routing
COPY ./static ./static


RUN go get -d -v
RUN cd /build && go build main.go

EXPOSE 8082
ENTRYPOINT ["/build/main"]