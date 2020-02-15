FROM golang:latest

WORKDIR /go/src/fizzimage
COPY . .

RUN go get -d -v .
RUN go install -v .

CMD fizzimg