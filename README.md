# go-statsd
[![Build Status](https://travis-ci.org/AstromechZA/go-statsd.svg)](https://travis-ci.org/AstromechZA/go-statsd) [![Code Coverage](https://gocover.io/_badge/github.com/AstromechZA/go-statsd)](https://gocover.io/github.com/AstromechZA/go-statsd) [![Documentation](https://godoc.org/github.com/AstromechZA/go-statsd?status.svg)](https://godoc.org/github.com/AstromechZA/go-statsd)

> Forked from https://github.com/alexcesaro/statsd on 2017-12-23 to add some missing bits. Mostly just fixing some of the filed issues, and merging some PR code. The upstream project is non-maintained and can be considered dead.
> For my own sake, I removed some of the gopkg branch anchors to reduce complexity.

## Introduction

statsd is a simple and efficient [Statsd](https://github.com/etsy/statsd)
client.

## Features

- Supports all StatsD metrics: counter, gauge, timing and set
- Supports InfluxDB and Datadog tags
- Fast and GC-friendly: all functions for sending metrics do not allocate
- Efficient: metrics are buffered by default
- Simple and clean API
- 100% test coverage
- An annoying hyphen in the name

## Usage

The basic usage can be very simple, see [example_test](./example_test.go). The documentation illustrates the more complex API methods and options.

## Documentation

https://godoc.org/github.com/AstromechZA/go-statsd

## Download

```
$ go get github.com/AstromechZA/go-statsd
```

## License

[MIT](LICENSE)
