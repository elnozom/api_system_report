package model

type models interface {
	req() GetBalanceOfTradeRequest
	resp() GetBalanceOfTradeResponse
}
