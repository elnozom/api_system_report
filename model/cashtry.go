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

type CashtryData struct {
	CashTryNo           int
	SessionNo           int
	EmpCode             int
	OpenDate            string
	CloseDate           string
	OpenTime            string
	CloseTime           string
	StartCash           float64
	TotalCash           float64
	ComputerName        string
	TotalOrder          float64
	TotalHome           float64
	TotalIn             float64
	TotalOut            float64
	TotalVisa           float64
	TotalShar           float64
	TotalVoid           float64
	Halek               float64
	EndCash             float64
	Paused              bool
	CasherMoney         float64
	PayLater            float64
	HomeIn              float64
	HomeOutCashTry      float64
	CashTryTypeCode     float64
	Final               float64
	CasherCashTrySerial float64
	Exceed              float64
	Shortage            float64
	StoreCode           float64
	TotalVat            float64
	DiscValue           float64
	TotalVoidCash       float64
	TotalVoidCrdt       float64
	TotalVoidVisa       float64
	DeliveryNonReturn   float64
}
