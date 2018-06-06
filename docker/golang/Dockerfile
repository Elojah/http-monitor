FROM golang:1.10

WORKDIR /go/src/app
COPY . .
RUN make dep
RUN make monitor
ENTRYPOINT /go/src/app/bin/http_monitor -c /go/src/app/bin/config.json

EXPOSE 5000
