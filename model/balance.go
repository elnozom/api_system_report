package model

type GetBalanceOfTradeRequest struct {
	AccType  int
	FromDate string
	ToDate   string
	PayCheq  bool
}

type GetBalanceOfTradeResponse struct {
	AccountCode int
	AccountName string
	AccNo       int
	BBC         float64
	BBD         float64
	BAC         float64
	BAD         float64
	Debit       float64
	Credit      float64
}
