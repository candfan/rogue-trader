
[![Build Status](https://travis-ci.org/evsamsonov/trading-indicators.svg?branch=master)](https://travis-ci.org/evsamsonov/trading-indicators)
[![Go Report Card](https://goreportcard.com/badge/github.com/evsamsonov/trading-indicators)](https://goreportcard.com/report/github.com/evsamsonov/trading-indicators)

# Trading indicators

The set of trading indicators

## Installation

```bash
$ go get github.com/evsamsonov/trading-indicators/indicator  
```

## Usage

All indicators requires trading data in [**timeseries**](https://github.com/evsamsonov/trading-timeseries) structure 

```go
dataset := []struct {