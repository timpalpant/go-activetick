package activetick

import (
	"time"
)

type SymbolStatus int

const (
	SymbolStatusSuccess      SymbolStatus = 1
	SymbolStatusInvalid      SymbolStatus = 2
	SymbolStatusUnavailable  SymbolStatus = 3
	SymbolStatusNoPermission SymbolStatus = 4
)

type QuoteField int

const (
	QuoteFieldSymbol                           QuoteField = 1
	QuoteFieldOpenPrice                        QuoteField = 2
	QuoteFieldPreviousClosePrice               QuoteField = 3
	QuoteFieldClosePrice                       QuoteField = 4
	QuoteFieldLastPrice                        QuoteField = 5
	QuoteFieldBidPrice                         QuoteField = 6
	QuoteFieldAskPrice                         QuoteField = 7
	QuoteFieldHighPrice                        QuoteField = 8
	QuoteFieldLowPrice                         QuoteField = 9
	QuoteFieldDayHighPrice                     QuoteField = 10
	QuoteFieldDayLowPrice                      QuoteField = 11
	QuoteFieldPreMarketOpenPrice               QuoteField = 12
	QuoteFieldExtendedHoursLastPrice           QuoteField = 13
	QuoteFieldAfterMarketClosePrice            QuoteField = 14
	QuoteFieldBidExchange                      QuoteField = 15
	QuoteFieldAskExchange                      QuoteField = 16
	QuoteFieldLastExchange                     QuoteField = 17
	QuoteFieldLastCondition                    QuoteField = 18
	QuoteFieldQuoteCondition                   QuoteField = 19
	QuoteFieldLastTradeDateTime                QuoteField = 20
	QuoteFieldLastQuoteDateTime                QuoteField = 21
	QuoteFieldDayHighDateTime                  QuoteField = 22
	QuoteFieldDayLowDateTime                   QuoteField = 23
	QuoteFieldLastSize                         QuoteField = 24
	QuoteFieldBidSize                          QuoteField = 25
	QuoteFieldAskSize                          QuoteField = 26
	QuoteFieldVolume                           QuoteField = 27
	QuoteFieldPreMarketVolume                  QuoteField = 28
	QuoteFieldAfterMarketVolume                QuoteField = 29
	QuoteFieldTradeCount                       QuoteField = 30
	QuoteFieldPreMarketTradeCount              QuoteField = 31
	QuoteFieldAfterMarketTradeCount            QuoteField = 32
	QuoteFieldFundamentalEquityName            QuoteField = 33
	QuoteFieldFundamentalEquityPrimaryExchange QuoteField = 34
)

type DataItemType int

const (
	DataByte          DataItemType = 1
	DataByteArray     DataItemType = 2
	DataUInteger32    DataItemType = 3
	DataUInteger64    DataItemType = 4
	DataInteger32     DataItemType = 5
	DataInteger64     DataItemType = 6
	DataPrice         DataItemType = 7
	DataString        DataItemType = 8
	DataUnicodeString DataItemType = 9
	DataDateTime      DataItemType = 10
	DataDouble        DataItemType = 11
)

type QuoteDataRequest struct {
	Symbols     []string
	QuoteFields []QuoteField
}

type QuoteDataResponse struct {
	Records []*QuoteSnapshotRecord
}

type QuoteSnapshotRecord struct {
	Symbol                           string
	OpenPrice                        float64
	PreviousClosePrice               float64
	ClosePrice                       float64
	LastPrice                        float64
	BidPrice                         float64
	AskPrice                         float64
	HighPrice                        float64
	LowPrice                         float64
	DayHighPrice                     float64
	DayLowPrice                      float64
	PreMarketOpenPrice               float64
	ExtendedHoursLastPrice           float64
	AfterMarketClosePrice            float64
	BidExchange                      Exchange
	AskExchange                      Exchange
	LastExchange                     Exchange
	LastCondition                    int
	QuoteCondition                   int
	LastTradeTime                    time.Time
	LastQuoteTime                    time.Time
	DayHighTime                      time.Time
	DayLowTime                       time.Time
	LastSize                         int
	BidSize                          int
	AskSize                          int
	Volume                           int
	PreMarketVolume                  int
	AfterMarketVolume                int
	TradeCount                       int
	PreMarketTradeCount              int
	AfterMarketTradeCount            int
	FundamentalEquityName            string
	FundamentalEquityPrimaryExchange Exchange
}

type QuoteStreamRequest struct {
	Symbols []string
}

type TradeStreamRecord struct {
	Symbol          string
	Flags           TradeFlag
	TradeConditions [4]TradeCondition
	LastExchange    Exchange
	LastPrice       float64
	LastSize        int
	LastDate        time.Time
}

type QuoteStreamRecord struct {
	Symbol         string
	QuoteCondition int
	BidExchange    Exchange
	AskExchange    Exchange
	BidPrice       float64
	AskPrice       float64
	BidSize        int
	AskSize        int
	QuoteTime      time.Time
}

type HistoryType int

const (
	HistoryTypeIntraday HistoryType = 0
	HistoryTypeDaily    HistoryType = 1
	HistoryTypeWeekly   HistoryType = 2
)

type BarDataRequest struct {
	Symbol          string
	HistoryType     HistoryType
	IntradayMinutes int
	BeginTime       time.Time
	EndTime         time.Time
}

type BarDataResponse struct {
	Records []*BarDataRecord
}

type BarDataRecord struct {
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

type TickDataRequest struct {
	Symbol    string
	Trades    bool
	Quotes    bool
	BeginTime time.Time
	EndTime   time.Time
}

type TickDataResponse struct {
	Records []*TickRecord
}

type TickType string

const (
	TickTypeQuote TickType = "Q"
	TickTypeTrade TickType = "T"
)

type TickRecord struct {
	Type         TickType
	Time         time.Time
	LastPrice    float64
	LastSize     int64
	LastExchange Exchange
	Condition    [4]TradeCondition
	BidPrice     float64
	AskPrice     float64
	BidSize      int64
	AskSize      int64
	BidExchange  Exchange
	AskExchange  Exchange
}

type OptionChainRequest struct {
	Symbol string
}

type OptionChainResponse struct {
	Records []string
}

type TradeFlag int

const (
	TradeFlagRegularMarketLastPrice  TradeFlag = 0x1
	TradeFlagRegularMarketVolume     TradeFlag = 0x2
	TradeFlagHighPrice               TradeFlag = 0x4
	TradeFlagLowPrice                TradeFlag = 0x8
	TradeFlagDayHighPrice            TradeFlag = 0x10
	TradeFlagDayLowPrice             TradeFlag = 0x20
	TradeFlagExtendedMarketLastPrice TradeFlag = 0x40
	TradeFlagPreMarketVolume         TradeFlag = 0x80
	TradeFlagAfterMarketVolume       TradeFlag = 0x100
	TradeFlagPreMarketOpenPrice      TradeFlag = 0x200
	TradeFlagOpenPrice               TradeFlag = 0x400
)

type TradeCondition int

const (
	TradeConditionRegular                       TradeCondition = 0
	TradeConditionAcquisition                   TradeCondition = 1
	TradeConditionAveragePrice                  TradeCondition = 2
	TradeConditionAutomaticExecution            TradeCondition = 3
	TradeConditionBunched                       TradeCondition = 4
	TradeConditionBunchSold                     TradeCondition = 5
	TradeConditionCAPElection                   TradeCondition = 6
	TradeConditionCash                          TradeCondition = 7
	TradeConditionClosing                       TradeCondition = 8
	TradeConditionCross                         TradeCondition = 9
	TradeConditionDerivativelyPriced            TradeCondition = 10
	TradeConditionDistribution                  TradeCondition = 11
	TradeConditionFormT                         TradeCondition = 12
	TradeConditionFormTOutOfSequence            TradeCondition = 13
	TradeConditionInterMarketSweep              TradeCondition = 14
	TradeConditionMarketCenterOfficialClose     TradeCondition = 15
	TradeConditionMarketCenterOfficialOpen      TradeCondition = 16
	TradeConditionMarketCenterOpening           TradeCondition = 17
	TradeConditionMarketCenterReOpenning        TradeCondition = 18
	TradeConditionMarketCenterClosing           TradeCondition = 19
	TradeConditionNextDay                       TradeCondition = 20
	TradeConditionPriceVariation                TradeCondition = 21
	TradeConditionPriorReferencePrice           TradeCondition = 22
	TradeConditionRule155Amex                   TradeCondition = 23
	TradeConditionRule127Nyse                   TradeCondition = 24
	TradeConditionOpening                       TradeCondition = 25
	TradeConditionOpened                        TradeCondition = 26
	TradeConditionRegularStoppedStock           TradeCondition = 27
	TradeConditionReOpening                     TradeCondition = 28
	TradeConditionSeller                        TradeCondition = 29
	TradeConditionSoldLast                      TradeCondition = 30
	TradeConditionSoldLastStoppedStock          TradeCondition = 31
	TradeConditionSoldOutOfSequence             TradeCondition = 32
	TradeConditionSoldOutOfSequenceStoppedStock TradeCondition = 33
	TradeConditionSplit                         TradeCondition = 34
	TradeConditionStockOption                   TradeCondition = 35
	TradeConditionYellowFlag                    TradeCondition = 36
)

type Exchange string

const (
	ExchangeAMEX                            Exchange = "A"
	ExchangeNasdaqOmxBx                     Exchange = "B"
	ExchangeNationalStockExchange           Exchange = "C"
	ExchangeFinraAdf                        Exchange = "D"
	ExchangeCQS                             Exchange = "E"
	ExchangeForex                           Exchange = "F"
	ExchangeInternationalSecuritiesExchange Exchange = "I"
	ExchangeEdgaExchange                    Exchange = "J"
	ExchangeEdgxExchange                    Exchange = "K"
	ExchangeChicagoStockExchange            Exchange = "M"
	ExchangeNyseEuronext                    Exchange = "N"
	ExchangeNyseArcaExchange                Exchange = "P"
	ExchangeNasdaqOmx                       Exchange = "Q"
	ExchangeCTS                             Exchange = "S"
	ExchangeCTANasdaqOMX                    Exchange = "T"
	ExchangeOTCBB                           Exchange = "U"
	ExchangeNNOTC                           Exchange = "u"
	ExchangeChicagoBoardOptionsExchange     Exchange = "W"
	ExchangeNasdaqOmxPhlx                   Exchange = "X"
	ExchangeBatsYExchange                   Exchange = "Y"
	ExchangeBatsExchange                    Exchange = "Z"
	ExchangeCanadaToronto                   Exchange = "T"
	ExchangeCanadaVenture                   Exchange = "V"
	ExchangeComposite                       Exchange = " "
)
