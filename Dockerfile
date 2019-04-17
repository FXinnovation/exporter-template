FROM golang:1.12 as builder
WORKDIR /go/src/github.com/FXinnovation/exporter-template
COPY . .
RUN make build

FROM quay.io/prometheus/busybox:glibc AS app
LABEL maintainer="FXinnovation CloudToolDevelopment <CloudToolDevelopment@fxinnovation.com>"
COPY --from=builder /go/src/github.com/FXinnovation/exporter-template/exporter-template /bin/exporter-template

EXPOSE      1984
WORKDIR /
ENTRYPOINT  [ "/bin/exporter-template"]
