FROM golang:1.16.3 AS tester
WORKDIR /app
COPY . .
RUN go get -d -v
RUN go test

FROM tester AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o app .

FROM alpine:latest
EXPOSE 80
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]