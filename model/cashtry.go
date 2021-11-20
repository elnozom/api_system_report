package model

type Cashtry struct {
	TotalCash   float64
	TotalOrder  int
	TVisa       float64
	TVoid       float64
	MonthNo     int
	AverageCash float64
	NoOfCashTry int
	AvgBasket   float64
}

type PausedCashtry struct {
	Serial       int
	EmpCode      int
	EmpName      string
	OpenDate     string
	OpenTime     string
	ComputerName string
}

type OpenCashtryReq struct {
	Store uint `json:"store"`
	Year  uint `json:"year"`
}

type OpenCashtry struct {
	EmpCode       int
	OpenDate      string
	StartCash     int
	TotalCash     float64
	CompouterName string
	TotalOrder    int
	TotalVisa     float64
	StoreName     string
}

type CashtryReq struct {
	Store uint `json:"store"`
	Year  uint `json:"year"`
}
type CashtryStores struct {
	StoreCode int    `json:"store_code"`
	StoreName string `json:"store_name"`
}
