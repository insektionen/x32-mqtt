FROM golang:1.18.0 AS builder
COPY . /app
RUN cd /app && go mod verify && go build

FROM debian:bullseye
COPY --from=builder /app/x32-mqtt /usr/bin/x32-mqtt

CMD ["/usr/bin/x32-mqtt"]
