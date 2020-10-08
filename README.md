
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
    Time   time.Time
    High   float64
    Low    float64
    Open   float64
    Close  float64
    Volume int64
}{
    {Time: time.Unix(1121979600, 0), High: 23, Low: 21.27, Open: 21.3125, Close: 22.1044, Volume: 4604900},
    {Time: time.Unix(1122238800, 0), High: 23.31999, Low: 22.15, Open: 22.15, Close: 23.21608, Volume: 4132600},
}

series := timeseries.New()
for _, item := range dataset {
    candle := timeseries.NewCandle(item.Time)