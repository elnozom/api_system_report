package model

type Topsale struct {
	ItemName    string
	TotalAmount float64
	TotalQnt    float64
	Profit      float64
	Disc        float64
}

type TopsaleReq struct {
	Store uint
	Date  string
}
