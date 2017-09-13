package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/timpalpant/go-activetick"
)

func fetchBarData(client *activetick.Client, symbol string, start, end time.Time) {
	req := &activetick.BarDataRequest{
		Symbol:          symbol,
		HistoryType:     activetick.HistoryTypeIntraday,
		IntradayMinutes: 1,
		BeginTime:       start,
		EndTime:         end,
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

func fetchTickData(client *activetick.Client, symbol string, start, end time.Time) {
	req := &activetick.TickDataRequest{
		Symbol:    symbol,
		BeginTime: start,
		EndTime:   end,
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

func main() {
	host := flag.String("host", "localhost", "ActiveTick HTTP server host")
	port := flag.Int("port", 5000, "ActiveTick HTTP port")
	symbol := flag.String("symbol", "SPY", "Symbol to fetch data for")
	dataType := flag.String("type", "bar", "Type of data to fetch (tick/bar)")
	beginTime := flag.String("begin_time", "2016-10-04T14:30:00Z", "Earliest time to fetch (RFC3339)")
	endTime := flag.String("end_time", "2016-10-04T14:40:00Z", "Latest time to fetch (RFC3339)")
	flag.Parse()

	startDate, err := time.Parse(time.RFC3339, *beginTime)
	if err != nil {
		log.Fatal(err)
	}
	endDate, err := time.Parse(time.RFC3339, *endTime)
	if err != nil {
		log.Fatal(err)
	}

	endpoint := fmt.Sprintf("http://%s:%d", *host, *port)
	client := activetick.NewClient(&http.Client{}, endpoint)

	switch *dataType {
	case "bar":
		fetchBarData(client, *symbol, startDate, endDate)
	case "tick":
		fetchTickData(client, *symbol, startDate, endDate)
	default:
		log.Fatal("Invalid data type: %v", *dataType)
	}
}
