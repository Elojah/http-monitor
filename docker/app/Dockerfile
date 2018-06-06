FROM golang:1.10

WORKDIR /go/src/monitor
COPY . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN make dep
RUN make monitor
ENTRYPOINT bin/http_monitor
CMD ["-c", "/go/src/app/bin/config.json"]
