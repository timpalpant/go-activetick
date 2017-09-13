package activetick

import (
	"time"
)

const (
	maxBars  = 20000
	maxTicks = 100000
)

// PagingClient wraps Client to automatically manage
// multiple paged requests.
type PagingClient struct {
	client *Client
}

func NewPagingClient(client *Client) *PagingClient {
	return &PagingClient{client}
}

// The HTTP API will return at most 20,000 records.
// If there are more than that, only the latest 20,000 in time are returned.
// So to fetch all data we need to page backward.
func (pc *PagingClient) GetBarData(req *BarDataRequest) (*BarDataResponse, error) {
	resp := &BarDataResponse{}

	for {
		page, err := pc.client.GetBarData(req)
		if err != nil {
			return nil, err
		}

		resp.Records = append(page.Records, resp.Records...)
		oldestTime := resp.Records[0].Time
		if len(page.Records) < maxBars || !oldestTime.Before(req.EndTime) {
			break
		}

		req = &BarDataRequest{
			Symbol:          req.Symbol,
			HistoryType:     req.HistoryType,
			IntradayMinutes: req.IntradayMinutes,
			BeginTime:       req.BeginTime,
			EndTime:         oldestTime.Add(-time.Minute),
		}
	}

	return resp, nil
}

// The HTTP API will return at most 100,000 ticks.
// If there are more than that, only the first 100,000 in time are returned.
// So to fetch all data we need to page forward.
func (pc *PagingClient) GetTickData(req *TickDataRequest) (*TickDataResponse, error) {
	resp := &TickDataResponse{}

	for {
		page, err := pc.client.GetTickData(req)
		if err != nil {
			return nil, err
		}

		resp.Records = append(resp.Records, page.Records...)

		latestTime := page.Records[len(page.Records)-1].Time
		latestTime = latestTime.Truncate(time.Second)
		if len(page.Records) < maxTicks || !latestTime.After(req.BeginTime) {
			break
		}

		for i := len(resp.Records) - 1; i >= 0; i-- {
			t := resp.Records[i].Time
			if !t.Before(latestTime) {
				resp.Records = resp.Records[:i]
			} else {
				break
			}
		}

		req = &TickDataRequest{
			Symbol:    req.Symbol,
			Trades:    req.Trades,
			Quotes:    req.Quotes,
			BeginTime: latestTime,
			EndTime:   req.EndTime,
		}
	}

	return resp, nil
}
