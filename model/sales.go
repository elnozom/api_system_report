package model

type MonthlySales struct {
	Totalamount float64
	DocMonth    int
}

type MonthlySalesReq struct {
	Year uint `json:"Year"`
}
type RevenueResp struct {
	Docdate     uint
	Totalamount float64
	Profit      float64
}
type RevenueReq struct {
	Year  uint
	Month uint
	Day   uint
}

type DailtlySales struct {
	Totalamount float64
	DocDay      int
}

type DailySalesReq struct {
	Year  uint `json:"Year"`
	Month uint `json:"Month"`
}
