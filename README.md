# statsd
[![Build Status](https://travis-ci.org/AstromechZA/statsd.svg)](https://travis-ci.org/AstromechZA/statsd) [![Code Coverage](http://gocover.io/_badge/github.com/AstromechZA/statsd)](http://gocover.io/github.com/AstromechZA/statsd) [![Documentation](https://godoc.org/github.com/AstromechZA/statsd?status.svg)](https://godoc.org/github.com/AstromechZA/statsd)

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

## Documentation

https://godoc.org/github.com/AstromechZA/statsd

## Download

    go get github.com/AstromechZA/statsd

## License

[MIT](LICENSE)
