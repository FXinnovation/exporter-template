# exporter-template
[![Build Status](https://travis-ci.org/FXinnovation/exporter-template.svg?branch=master)](https://travis-ci.org/FXinnovation/exporter-template)

The exporter-template is a basic Prometheus exporter, written in Go. 
The goal of this project is to provide a minimal start point from which any exporter may start.
It contains a minimal set of features common to any exporter.

## Getting Started

In summary, copy this project's code to get your exporter started.

### Prerequisites

To run this project, you will need a [working Go environment](https://golang.org/doc/install).

### Installing

```bash
go get -u github.com/FXinnovation/exporter-template
```

## Running the tests

```bash
make test
```

## Usage

```bash
./exporter-template -h
```

## Deployment

The template listen on port 1984 by default, be sure to change that port in order to suit your new exporter needs.

## Contributing

Refer to [CONTRIBUTING.md](https://github.com/FXinnovation/exporter-template/blob/master/CONTRIBUTING.md).

## License

Apache License 2.0, see [LICENSE](https://github.com/FXinnovation/exporter-template/blob/master/LICENSE).
