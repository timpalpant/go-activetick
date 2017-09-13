package activetick

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	timeFormat = "20060102150405"
)

// Client provides methods to interact with the ActiveTick HTTP API.
// TODO: Implement snapshot quotes (/quoteData) and streaming quotes
// (/quoteStream).
type Client struct {
	client   *http.Client
	endpoint string
}

func NewClient(client *http.Client, endpoint string) *Client {
	return &Client{client, endpoint}
}

func (c *Client) GetBarData(req *BarDataRequest) (*BarDataResponse, error) {
	values := url.Values{}
	values.Set("symbol", req.Symbol)
	values.Set("historyType", strconv.Itoa(int(req.HistoryType)))
	values.Set("intradayMinutes", strconv.Itoa(req.IntradayMinutes))
	values.Set("beginTime", req.BeginTime.Format(timeFormat))
	values.Set("endTime", req.EndTime.Format(timeFormat))

	result, err := c.getCSV("/barData", values)
	if err != nil {
		return nil, err
	}

	resp := &BarDataResponse{
		Records: make([]*BarDataRecord, 0, len(result)),
	}

	for _, row := range result {
		record, err := parseBarData(row)
		if err != nil {
			return nil, err
		}

		resp.Records = append(resp.Records, record)
	}

	return resp, err
}

func parseBarData(row []string) (*BarDataRecord, error) {
	if len(row) != 6 {
		return nil, fmt.Errorf("Expected %d rows, got %d: %v",
			6, len(row), row)
	}

	t, err := time.Parse(timeFormat, row[0])
	if err != nil {
		return nil, err
	}

	open, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return nil, err
	}

	high, err := strconv.ParseFloat(row[2], 64)
	if err != nil {
		return nil, err
	}

	low, err := strconv.ParseFloat(row[3], 64)
	if err != nil {
		return nil, err
	}

	cl, err := strconv.ParseFloat(row[4], 64)
	if err != nil {
		return nil, err
	}

	vol, err := strconv.ParseInt(row[5], 10, 64)
	if err != nil {
		return nil, err
	}

	return &BarDataRecord{
		Time:   t,
		Open:   open,
		High:   high,
		Low:    low,
		Close:  cl,
		Volume: vol,
	}, nil
}

func (c *Client) GetTickData(req *TickDataRequest) (*TickDataResponse, error) {
	values := url.Values{}
	values.Set("symbol", req.Symbol)
	tradesFlag := "0"
	if req.Trades {
		tradesFlag = "1"
	}
	values.Set("trades", tradesFlag)
	quotesFlag := "0"
	if req.Quotes {
		quotesFlag = "1"
	}
	values.Set("quotes", quotesFlag)
	// NOTE: Milliseconds are not supported as suggested in the documentation.
	values.Set("beginTime", req.BeginTime.Format(timeFormat))
	values.Set("endTime", req.EndTime.Format(timeFormat))

	result, err := c.getCSV("/tickData", values)
	if err != nil {
		return nil, err
	}

	resp := &TickDataResponse{
		Records: make([]*TickRecord, 0, len(result)),
	}

	for _, row := range result {
		record, err := parseTickData(row)
		if err != nil {
			return nil, err
		}

		resp.Records = append(resp.Records, record)
	}

	return resp, err
}

func parseTickData(row []string) (*TickRecord, error) {
	if len(row) < 9 {
		return nil, fmt.Errorf("Expected %d rows, got %d: %v",
			9, len(row), row)
	}

	tickType := TickType(row[0])
	t, err := parseTime(row[1])
	if err != nil {
		return nil, err
	}

	record := &TickRecord{
		Type: tickType,
		Time: t,
	}

	switch tickType {
	case TickTypeTrade:
		price, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, err
		}

		size, err := strconv.ParseInt(row[3], 10, 64)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(record.Condition); i++ {
			tc, err := strconv.ParseInt(row[i+5], 10, 64)
			if err != nil {
				return nil, err
			}

			record.Condition[i] = TradeCondition(tc)
		}

		record.LastPrice = price
		record.LastSize = size
		record.LastExchange = Exchange(row[4])
	case TickTypeQuote:
		bidPrice, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			return nil, err
		}

		askPrice, err := strconv.ParseFloat(row[3], 64)
		if err != nil {
			return nil, err
		}

		bidSize, err := strconv.ParseInt(row[4], 10, 64)
		if err != nil {
			return nil, err
		}

		askSize, err := strconv.ParseInt(row[5], 10, 64)
		if err != nil {
			return nil, err
		}

		cond, err := strconv.ParseInt(row[8], 10, 64)
		if err != nil {
			return nil, err
		}

		record.BidPrice = bidPrice
		record.AskPrice = askPrice
		record.BidSize = bidSize
		record.AskSize = askSize
		record.BidExchange = Exchange(row[6])
		record.AskExchange = Exchange(row[7])
		record.Condition[0] = TradeCondition(cond)
	default:
		return nil, fmt.Errorf("Unknown tick type: %v", tickType)
	}

	return record, nil
}

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse(timeFormat, s[:len(s)-3])
	if err != nil {
		return t, err
	}

	ms, err := strconv.ParseInt(s[len(s)-3:], 10, 64)
	if err != nil {
		return t, err
	}

	return t.Add(time.Duration(ms) * time.Millisecond), nil
}

func (c *Client) getCSV(route string, values url.Values) ([][]string, error) {
	url := c.endpoint + route
	params := values.Encode()
	if params != "" {
		url = url + "?" + params
	}

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %v", resp.Status, string(body))
	}

	reader := csv.NewReader(resp.Body)
	return reader.ReadAll()
}
