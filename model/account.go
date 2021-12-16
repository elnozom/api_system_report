package model

type GetAccountRequest struct {
	Code int    `json:"Code" validate:"required"`
	Name string `json:"Name" validate:"required"`
	Type int    `json:"Type" validate:"required"`
}

type InsertStktr01Request struct {
	TransactionType int
	Safe            int
	AccSerial       int
	Amount          float64
	AccType         int
	Store           int
}

type GetAccountBalanceRequest struct {
	FromDate  string `validate:"required"`
	ToDate    string `validate:"required"`
	AccSerial int    `validate:"required"`
}

type GetAccountBalanceData struct {
	DocNo       string
	AccMoveName string
	DocDate     string
	Dbt         float64
	Crdt        float64
	RaseedCrdt  float64
	RaseedDbt   float64
}

type GetAccountBalanceResponse struct {
	Raseed float64
	Data   []GetAccountBalanceData
}

type Account struct {
	Serial      int
	AccountCode int
	AccountName string
}

type TransCycleAccReq struct {
	DateFrom      string
	DateTo        string
	StoreCode     int
	GroupCode     int
	AccountSerial int
}

type TransCycleAccResp struct {
	BuyWhole          float64
	BuyPart           float64
	SaleWhole         float64
	SalePart          float64
	TransOutWhole     float64
	TransOutPart      float64
	TransInWhole      float64
	TransInPart       float64
	IndusInWhole      float64
	IndusInPart       float64
	IndusOutWhole     float64
	IndusOutPart      float64
	RaseedbeforePart  float64
	RaseedbeforeWhole float64
	RaseedPart        float64
	RaseedWhole       float64
	CycleRate         float64
	LastBuyDate       string
	LastSellDate      string
	ItemName          string
	ItemCode          int
	GroupCode         int
	AccountSerial     int
	MinorPerMajor     int
	ByWeight          bool
}

type TransCycleAccRespHelper struct {
	Buy           float64
	Sale          float64
	TransOut      float64
	TransIn       float64
	IndusIn       float64
	IndusOut      float64
	Raseedbefore  float64
	Raseed        float64
	MinorPerMajor int
}
