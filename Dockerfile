FROM golang:1.18.6-alpine3.16 as build

RUN mkdir /build/
COPY . /build/
WORKDIR /build/

RUN GOOS=linux go build -o appointment-schedule ./cmd/main.go

FROM alpine as assemble
WORKDIR /opt/future/
COPY --from=build /build/appointment-schedule /opt/future/
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip
EXPOSE 8080
ENTRYPOINT ["./appointment-schedule"]
