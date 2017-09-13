# go-activetick
A Go library for accessing the ActiveTick HTTP API.

[![GoDoc](https://godoc.org/github.com/timpalpant/go-activetick?status.svg)](http://godoc.org/github.com/timpalpant/go-activetick)
[![Build Status](https://travis-ci.org/timpalpant/go-activetick.svg?branch=master)](https://travis-ci.org/timpalpant/go-activetick)
[![Coverage Status](https://coveralls.io/repos/timpalpant/go-activetick/badge.svg?branch=master&service=github)](https://coveralls.io/github/timpalpant/go-activetick?branch=master)

go-activetick is a library to access the [ActiveTick](http://www.activetick.com/activetick/contents/PersonalServicesDataAPIOverview.aspx) HTTP API from [Go](http://www.golang.org).

ActiveTick is not affiliated and does not endorse or recommend this library.

## Usage

### atclient CLI

```
$ atclient -symbol SPY -begin_time 2016-10-04T14:30:00Z -end_time 2016-10-04T14:40:00Z -type tick
```

### Fetch historical minute bars

```Go
package main

import (
    "fmt"
    "net/http"
    "time"

    "github.com/timpalpant/go-activetick"
)

func main() {
    endpoint := "http://localhost:5000"
    client := activetick.NewPagingClient(activetick.NewClient(&http.Client{}, endpoint))

	req := &activetick.BarDataRequest{
		Symbol:          "SPY",
		HistoryType:     activetick.HistoryTypeIntraday,
		IntradayMinutes: 1,
		BeginTime:       time.Now().Add(-48 * time.Hour),
		EndTime:         time.Now().Add(-47 * time.Hour),
	}

	resp, err := client.GetBarData(req)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range resp.Records {
		fmt.Printf("%v,%v,%v,%v,%v,%v\n", record.Time, record.Open,
			record.High, record.Low, record.Close, record.Volume)
	}
}
```

### Fetch historical ticks

```Go
package main

import (
    "fmt"
    "net/http"

    "github.com/timpalpant/go-activetick"
)

func main() {
    endpoint := "http://localhost:5000"
    client := activetick.NewPagingClient(activetick.NewClient(&http.Client{}, endpoint))

	req := &activetick.TickDataRequest{
		Symbol:    "SPY",
		BeginTime: time.Now().Add(-48 * time.Hour),
		EndTime:   time.Now().Add(-47 * time.Hour),
		Trades:    true,
		Quotes:    true,
	}

	resp, err := client.GetTickData(req)
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range resp.Records {
		if record.Type == activetick.TickTypeQuote {
			fmt.Printf("%v,%v,%v,%v,%v,%v,%v\n", record.Time,
				record.BidPrice, record.BidSize, record.BidExchange,
				record.AskPrice, record.AskSize, record.AskExchange)
		} else {
			fmt.Printf("%v,%v,%v,%v\n", record.Time,
				record.LastPrice, record.LastSize, record.LastExchange)
		}
	}
}
```

## Contributing

Pull requests and issues are welcomed!

## License

go-activetick is released under the [GNU Lesser General Public License, Version 3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html)
