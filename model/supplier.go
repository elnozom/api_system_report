package model

type Supplier struct {
	AccountCode    int
	AccountName    string
	DBT            float32
	CRDT           float32
	ReturnBuy      float32
	Buy            float32
	Paid           float32
	CHEQUE         float32
	CHQUnderCollec float32
	Discount       float32
}
type SupplierBalance struct {
	AccountCode    int
	AccountName    string
	DBT            float64
	CRDT           float64
	ReturnBuy      float64
	Buy            float64
	Paid           float64
	CHEQUE         float64
	CHQUnderCollec float64
	Discount       float64
	NetBuy         float64
	BalanceDebit   float64
	BalanceCredit  float64
}
