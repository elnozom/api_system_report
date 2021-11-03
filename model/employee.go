package model

type EmpTotalsReq struct {
	EmpCode  int
	FromDate string `validate:"required"`
	ToDate   string `validate:"required"`
}

type EmpTotalsResp struct {
	Orders  int
	Amount  float64
	ROrders float64
	RAmount float64
	EmpCode int
	EmpName string
}

type Emp struct {
	EmpName string
	EmpCode int
}
