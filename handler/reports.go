package handler

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/ahmedashrafdev/reports/db"
	"github.com/ahmedashrafdev/reports/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CashTryAnalysis(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.CashtryReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var cashtries []model.Cashtry
	rows, err := db.Raw("EXEC CashtryAnalysis @StoreCode = ?, @Year = ?;", req.Store, req.Year).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cashtry model.Cashtry
		rows.Scan(&cashtry.TotalCash, &cashtry.TotalOrder, &cashtry.TVisa, &cashtry.TVoid, &cashtry.MonthNo, &cashtry.AverageCash, &cashtry.NoOfCashTry, &cashtry.AvgBasket)
		cashtries = append(cashtries, cashtry)
	}

	return c.JSON(http.StatusOK, cashtries)
}

func (h *Handler) GetDrivers(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	var employee []model.Emp
	rows, err := db.Raw("EXEC EmployeeDriverList").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Emp
		err = rows.Scan(&item.EmpCode, &item.EmpName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		employee = append(employee, item)
	}

	return c.JSON(http.StatusOK, employee)
}

func (h *Handler) EmpTotals(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.EmpTotalsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var resp []model.EmpTotalsResp
	rows, err := db.Raw("EXEC GetDriverTotals @Empcode = ?, @DateFrom = ? ,@Dateto = ?;", req.EmpCode, req.FromDate, req.ToDate).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.EmpTotalsResp
		rows.Scan(&item.Orders, &item.Amount, &item.ROrders, &item.RAmount, &item.EmpCode, &item.EmpName)
		resp = append(resp, item)
	}

	return c.JSON(http.StatusOK, resp)
}
func (h *Handler) GetRevenuePerTime(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server"+err.Error())
	}
	db := db.DBConn
	req := new(model.RevenueReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var resp []model.RevenueResp
	rows, err := db.Raw("EXEC Rpt_revenueProfitPerTime @Year = ?, @Month = ? ,@Day = ?;", req.Year, req.Month, req.Day).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.RevenueResp
		rows.Scan(&item.Docdate, &item.Totalamount, &item.Profit)
		resp = append(resp, item)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetExpnsesByMonth(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server"+err.Error())
	}
	db := db.DBConn
	req := new(model.ExpensesReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var resp []model.ExpensesResp
	rows, err := db.Raw("EXEC Rpt_Expenses @Year = ?;", req.Year).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.ExpensesResp
		rows.Scan(&item.Docdate, &item.Totalamount)
		resp = append(resp, item)
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) AccTr01Insert(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server"+err.Error())
	}
	db := db.DBConn
	req := new(model.InsertStktr01Request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request"+err.Error())
	}
	print(req)
	_, err = db.Raw(
		"EXEC AccTr01Insert  @AccMoveSerial = ? , @Stc = ? , @UserCode = ?, @AccType = ? ,@AccountSerial = ? ,@AccountSerial2 = ? ,@Amount =? ; ", req.TransactionType, req.Store, 1, req.AccType, req.Safe, req.AccSerial, req.Amount).Rows()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}
func (h *Handler) GetAccountBalance(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}

	db := db.DBConn
	req := new(model.GetAccountBalanceRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var resp model.GetAccountBalanceResponse
	var data []model.GetAccountBalanceData
	dateRows, err := db.Raw("EXEC AccTr01GetBalancBefore @DateFrom = ?, @AccSerial = ? ;", req.FromDate, req.AccSerial).Rows()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	rows, err := db.Raw("EXEC AccTr01CashFlow @DateFrom = ?, @DateTo = ? , @AccSerial = ? ;", req.FromDate, req.ToDate, req.AccSerial).Rows()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer rows.Close()
	defer dateRows.Close()

	for dateRows.Next() {
		dateRows.Scan(&resp.Raseed)
	}
	var raseed = math.Abs(resp.Raseed)
	for rows.Next() {
		var rec model.GetAccountBalanceData
		rows.Scan(&rec.DocNo, &rec.DocDate, &rec.Dbt, &rec.Crdt)
		r := raseed + (rec.Crdt - rec.Dbt)
		if r > 0 {
			rec.RaseedCrdt = r
		} else {
			rec.RaseedDbt = math.Abs(r)
		}
		raseed = math.Abs(r)
		data = append(data, rec)
	}
	resp.Data = data

	return c.JSON(http.StatusOK, resp.Data)
}

func (h *Handler) GetDocNo(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var DocNo []model.Doc
	rows, err := db.Raw("EXEC GetSdDocNo @DevNo = ?, @TrSerial = ?,@StoreCode = ?;", req.DevNo, req.TrSerial, req.StoreCode).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var doc model.Doc
		err = rows.Scan(
			&doc.DocNo,
		)
		print(rows)
		if err != nil {
			return c.JSON(http.StatusOK, 1)
		}
		DocNo = append(DocNo, doc)
	}

	return c.JSON(http.StatusOK, DocNo[0].DocNo+1)
}

func (h *Handler) GetOpenDocs(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.OpenDocReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var OpenDocs []model.OpenDoc
	rows, err := db.Raw("EXEC GetOpenSdDocNo @DevNo = ?, @TrSerial = ?;", req.DevNo, req.TrSerial).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var openDoc model.OpenDoc
		err = rows.Scan(
			&openDoc.DocNo,
			&openDoc.StoreCode,
			&openDoc.AccontSerial,
			&openDoc.TransSerial,
			&openDoc.AccountName,
			&openDoc.AccountCode,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		OpenDocs = append(OpenDocs, openDoc)
	}

	return c.JSON(http.StatusOK, OpenDocs)
}
func (h *Handler) GetCashFlow(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.CashFlowReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var CashFlows []model.CashFlow
	rows, err := db.Raw("EXEC cashFlow @DateFrom = ?, @DateTo = ?,@AccSerial = ?;", req.FromDate, req.ToDate, req.AccSerial).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var CashFlow model.CashFlow
		err = rows.Scan(
			&CashFlow.DocDate,
			&CashFlow.Income,
			&CashFlow.Supplier,
			&CashFlow.Expensis,
			&CashFlow.Others,
			&CashFlow.Bankin,
			&CashFlow.Cheqout,
			&CashFlow.Cheqin,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		CashFlows = append(CashFlows, CashFlow)
	}

	// return c.JSON(http.StatusOK, "success")
	return c.JSON(http.StatusOK, CashFlows)

}

func (h *Handler) GetSupplierBalance(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	var Suppliers []model.SupplierBalance
	rows, err := db.Raw("EXEC GetSupplierBalance").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var Supplier model.SupplierBalance
		err = rows.Scan(
			&Supplier.AccountCode,
			&Supplier.AccountName,
			&Supplier.DBT,
			&Supplier.CRDT,
			&Supplier.ReturnBuy,
			&Supplier.Buy,
			&Supplier.Paid,
			&Supplier.CHEQUE,
			&Supplier.CHQUnderCollec,
			&Supplier.Discount,
		)

		Supplier.NetBuy = Supplier.Buy - Supplier.ReturnBuy - Supplier.Discount
		var Balance float64
		Balance = Supplier.NetBuy + Supplier.CRDT - Supplier.Paid - Supplier.CHEQUE - Supplier.DBT
		if Balance > 0 {
			Supplier.BalanceDebit = Balance
		} else {
			Supplier.BalanceCredit = Balance
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		Suppliers = append(Suppliers, Supplier)
	}

	// return c.JSON(http.StatusOK, "success")
	return c.JSON(http.StatusOK, Suppliers)

}
func (h *Handler) GetCashFlowYear(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.CashFlowYearReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var CashFlows []model.CashFlow
	rows, err := db.Raw("EXEC cashFlowYear @Year = ? ,@AccSerial = ?;", req.Year, req.AccSerial).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var CashFlow model.CashFlow
		err = rows.Scan(
			&CashFlow.DocDate,
			&CashFlow.Income,
			&CashFlow.Supplier,
			&CashFlow.Expensis,
			&CashFlow.Others,
			&CashFlow.Bankin,
			&CashFlow.Cheqout,
			&CashFlow.Cheqin,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		CashFlows = append(CashFlows, CashFlow)
	}

	// return c.JSON(http.StatusOK, "success")
	return c.JSON(http.StatusOK, CashFlows)

}

func (h *Handler) GetDocItems(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DocItemsReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	var DocItems []model.DocItem
	rows, err := db.Raw("EXEC GetSdItems @DevNo = ?, @TrSerial = ?,@StoreCode = ? , @DocNo = ?;", req.DevNo, req.TrSerial, req.StoreCode, req.DocNo).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var docItem model.DocItem
		err = rows.Scan(
			&docItem.Serial,
			&docItem.Qnt,
			&docItem.Item_BarCode,
			&docItem.ItemName,
			&docItem.MinorPerMajor,
			&docItem.ByWeight,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")
		}
		DocItems = append(DocItems, docItem)
	}

	return c.JSON(http.StatusOK, DocItems)
}

func (h *Handler) GetCashtrayData(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	var resp model.CashtryData
	err = db.Raw("EXEC CashTryData @CashTrySerial = ?;", c.Param("serial")).Row().Scan(
		&resp.CashTryNo,
		&resp.SessionNo,
		&resp.EmpCode,
		&resp.OpenDate,
		&resp.CloseDate,
		&resp.OpenTime,
		&resp.CloseTime,
		&resp.StartCash,
		&resp.TotalCash,
		&resp.ComputerName,
		&resp.TotalOrder,
		&resp.TotalHome,
		&resp.TotalIn,
		&resp.TotalOut,
		&resp.TotalVisa,
		&resp.TotalShar,
		&resp.TotalVoid,
		&resp.Halek,
		&resp.EndCash,
		&resp.Paused,
		&resp.CasherMoney,
		&resp.PayLater,
		&resp.HomeIn,
		&resp.HomeOutCashTry,
		&resp.CashTryTypeCode,
		&resp.Final,
		&resp.CasherCashTrySerial,
		&resp.Exceed,
		&resp.Shortage,
		&resp.StoreCode,
		&resp.TotalVat,
		&resp.DiscValue,
		&resp.TotalVoidCash,
		&resp.TotalVoidCrdt,
		&resp.TotalVoidVisa,
	)
	resp.DeliveryNonReturn = resp.TotalHome - resp.HomeIn
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DeleteItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := db.Raw("EXEC DeleteSdItem  @Serial = ?; ", req.Serial).Rows()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) InsertItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	print("asddd")

	db := db.DBConn
	req := new(model.InsertItemReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "ERROR binding request")
	}
	print(req)
	rows, err := db.Raw(
		"EXEC InsertSdDocNo  @DNo = ? ,@TrS = ? ,@AccS = ? ,@ItmS =?  ,@Qnt = ? ,@StCode = ? ,@InvNo = ? ,@ItmBarCode = ? ,@DevNo = ?,@StCode2 = ?; ", req.DNo, req.TrS, req.AccS, req.ItmS, req.Qnt, req.StCode, req.InvNo, req.ItmBarCode, req.DevNo, req.StCode2).Rows()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, rows)
}
func (h *Handler) OpenCashTry(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.OpenCashtryReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var openCashtries []model.OpenCashtry
	rows, err := db.Raw("EXEC GetOpenCashTry").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cashtry model.OpenCashtry
		err = rows.Scan(
			&cashtry.EmpCode,
			&cashtry.OpenDate,
			&cashtry.StartCash,
			&cashtry.TotalCash,
			&cashtry.CompouterName,
			&cashtry.TotalOrder,
			&cashtry.TotalVisa,
			&cashtry.StoreName,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values")

		}
		openCashtries = append(openCashtries, cashtry)
	}

	return c.JSON(http.StatusOK, openCashtries)
}

func (h *Handler) PausedCashTry(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	var pausedCashtries []model.PausedCashtry
	rows, err := db.Raw("EXEC CashTrayPaused").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var cashtry model.PausedCashtry
		err = rows.Scan(
			&cashtry.Serial,
			&cashtry.EmpCode,
			&cashtry.EmpName,
			&cashtry.OpenDate,
			&cashtry.OpenTime,
			&cashtry.ComputerName,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "can't scan the values"+err.Error())

		}
		pausedCashtries = append(pausedCashtries, cashtry)
	}

	return c.JSON(http.StatusOK, pausedCashtries)
}

func (h *Handler) CashTryStores(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	db := db.DBConn
	var stores []model.CashtryStores
	rows, err := db.Raw("EXEC GetStoreName").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var store model.CashtryStores
		rows.Scan(&store.StoreCode, &store.StoreName)
		stores = append(stores, store)
	}

	return c.JSON(http.StatusOK, stores)
}

func (h *Handler) GetTopSalesItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	db := db.DBConn
	req := new(model.TopsaleReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var topsales []model.Topsale
	parseDate := strings.Split(req.Date, "-")
	rows, err := db.Raw("EXEC GetTopSalesItem @Year = ?, @Month = ?,@StoreCode = ?;", parseDate[0], parseDate[1], req.Store).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var topsale model.Topsale
		rows.Scan(&topsale.ItemName, &topsale.TotalAmount, &topsale.TotalQnt, &topsale.Profit, &topsale.Disc)
		// println(topsale)
		topsales = append(topsales, topsale)
	}

	return c.JSON(http.StatusOK, topsales)
}

func (h *Handler) GetBranchesSales(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.BranchesSaleReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var branchesSales []model.BranchesSale
	rows, err := db.Raw("EXEC GetBranchesSales @Year = ?, @Month = ?", req.Year, req.Month).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var branchSale model.BranchesSale
		rows.Scan(&branchSale.StoreCode, &branchSale.StoreName, &branchSale.Totalamount)
		// println(topsale)
		branchesSales = append(branchesSales, branchSale)
	}

	return c.JSON(http.StatusOK, branchesSales)
}

func (h *Handler) GetBranchesProfit(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.BranchesProfitReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var branchesProfit []model.BranchesProfit
	rows, err := db.Raw("EXEC GetBranchesProfit @Year = ?, @Month = ?", req.Year, req.Month).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var branchSale model.BranchesProfit
		rows.Scan(&branchSale.StoreCode, &branchSale.StoreName, &branchSale.Profit)
		branchesProfit = append(branchesProfit, branchSale)
	}

	return c.JSON(http.StatusOK, branchesProfit)
}

func (h *Handler) GetAccount(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	fmt.Println(req)

	var accounts []model.Account
	rows, err := db.Raw("EXEC GetAccount @Code = ?, @Name = ? , @Type = ?", req.Code, req.Name, req.Type).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var account model.Account
		rows.Scan(&account.Serial, &account.AccountCode, &account.AccountName)
		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

func (h *Handler) GetItem(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.GetItemRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var rows *sql.Rows
	var rowsErr error
	var items []model.Item
	if req.Name == "" {
		rows, rowsErr = db.Raw("EXEC GetItemData @BCode = ?, @StoreCode = ? ", req.BCode, req.StoreCode).Rows()
	} else {
		rows, rowsErr = db.Raw("EXEC GetItemData @BCode = ?, @StoreCode = ? , @Name = ? ", req.BCode, req.StoreCode, req.Name).Rows()
	}
	if rowsErr != nil {
		return c.JSON(http.StatusInternalServerError, rowsErr.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var item model.Item
		err = rows.Scan(&item.Serial, &item.ItemName, &item.MinorPerMajor, &item.POSPP, &item.POSTP, &item.ByWeight, &item.WithExp, &item.ItemHasAntherUnit, &item.AvrWait, &item.Expirey, &item.I, &item.R)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}
func (h *Handler) GetMonthlySales(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.MonthlySalesReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var monthlySales []model.MonthlySales
	rows, err := db.Raw("EXEC GetMonthlySales @Year = ?", req.Year).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err doing stored procedure")
	}

	defer rows.Close()
	for rows.Next() {
		var monthlySale model.MonthlySales
		if err := rows.Scan(&monthlySale.DocMonth, &monthlySale.Totalamount); err != nil {
			panic(err)
		}
		fmt.Printf("is %x", monthlySale.DocMonth)
		// println(topsale)
		monthlySales = append(monthlySales, monthlySale)
	}

	return c.JSON(http.StatusOK, monthlySales)
}

func (h *Handler) GetDailySales(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.DailySalesReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var dailySales []model.DailtlySales
	rows, err := db.Raw("EXEC GetDailySales @Month = ? , @Year = ?", req.Month, req.Year).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err doing stored procedure")
	}

	defer rows.Close()
	for rows.Next() {
		var monthlySale model.DailtlySales
		if err := rows.Scan(&monthlySale.DocDay, &monthlySale.Totalamount); err != nil {
			panic(err)
		}
		fmt.Printf("is %x", monthlySale.DocDay)
		// println(topsale)
		dailySales = append(dailySales, monthlySale)
	}

	return c.JSON(http.StatusOK, dailySales)
}

func (h *Handler) GetBalnaceOfTrade(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.GetBalanceOfTradeRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp []model.GetBalanceOfTradeResponse
	var rows *sql.Rows
	var rowErr error
	if req.PayCheq {
		rows, rowErr = db.Raw("EXEC balanceoftrade1 @AccountType = ? , @DateFrom = ? , @DateTo = ?", req.AccType, req.FromDate, req.ToDate).Rows()
	} else {
		rows, rowErr = db.Raw("EXEC balanceoftrade @AccountType = ? , @DateFrom = ? , @DateTo = ?", req.AccType, req.FromDate, req.ToDate).Rows()
	}
	if rowErr != nil {
		return c.JSON(http.StatusInternalServerError, "err doing stored procedure"+rowErr.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var rec model.GetBalanceOfTradeResponse
		if err := rows.Scan(&rec.AccountCode, &rec.AccountName, &rec.AccNo, &rec.BBC, &rec.BBD, &rec.BAC, &rec.BAD); err != nil {
			return c.JSON(http.StatusInternalServerError, "err scanning result"+err.Error())
		}

		resp = append(resp, rec)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetTransCycleAcc(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to connect to your server")
	}
	db := db.DBConn
	req := new(model.TransCycleAccReq)
	if err := c.Bind(req); err != nil {
		return err
	}
	var resp []model.TransCycleAccResp
	var rows *sql.Rows
	var rowErr error
	rows, rowErr = db.Raw("EXEC Rpt_TransCycleAcc  @DateFrom = ? , @DateTo = ? , @Storeode = ? , @GroupCode = ? , @AccSerial = ?", req.DateFrom, req.DateTo, req.StoreCode, req.GroupCode, req.AccountSerial).Rows()

	if rowErr != nil {
		return c.JSON(http.StatusInternalServerError, "err doing stored procedure "+rowErr.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var rec model.TransCycleAccResp
		var helper model.TransCycleAccRespHelper
		if err := rows.Scan(&helper.Buy, &helper.Sale, &helper.TransOut, &helper.TransIn, &helper.IndusIn, &helper.IndusOut, &helper.Raseedbefore, &helper.Raseed, &rec.LastBuyDate, &rec.LastSellDate, &rec.ItemName, &rec.ItemCode, &rec.GroupCode, &rec.AccountSerial, &helper.MinorPerMajor, &rec.ByWeight); err != nil {
			return c.JSON(http.StatusInternalServerError, "err scanning result"+err.Error())
		}
		calculateCycle(&helper, &rec)

		resp = append(resp, rec)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetGroups(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	db := db.DBConn
	var groups []model.Group
	rows, err := db.Raw("EXEC GroupCodeListAll").Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err "+err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var group model.Group
		rows.Scan(&group.GroupCode, &group.GroupName)
		groups = append(groups, group)
	}

	return c.JSON(http.StatusOK, groups)
}
func (h *Handler) GetStock(c echo.Context) error {
	err := h.userStore.ConnectDb(userIDFromToken(c))
	db := db.DBConn
	req := new(model.StockReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, "err scanning result"+err.Error())
	}
	var stock []model.StockResp

	rows, err := db.Raw("EXEC Rpt_Stock @StoreCode = ? , @GroupCode = ? , @ItemSerial = ?", req.Store, req.Group, req.Item).Rows()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err executing procedure "+err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println(req)
		var item model.StockResp
		rows.Scan(&item.ItemCode, &item.ItemName, &item.Raseed)
		stock = append(stock, item)
	}

	return c.JSON(http.StatusOK, stock)
}

func calculateCycle(helper *model.TransCycleAccRespHelper, rec *model.TransCycleAccResp) {
	rec.MinorPerMajor = helper.MinorPerMajor
	rightHand := helper.Raseedbefore + helper.Buy + helper.IndusIn + helper.TransIn
	leftHand := helper.Sale + helper.TransOut + helper.IndusOut

	if rightHand == 0 {
		rec.CycleRate = 0
	} else {
		rec.CycleRate = math.Floor((leftHand / rightHand) * 100)
	}

	rec.BuyPart, rec.BuyWhole = convertToPartAndWhole(helper.Buy, helper.MinorPerMajor)
	rec.SalePart, rec.SaleWhole = convertToPartAndWhole(helper.Sale, helper.MinorPerMajor)
	rec.TransInPart, rec.TransInWhole = convertToPartAndWhole(helper.TransIn, helper.MinorPerMajor)
	rec.TransOutPart, rec.TransOutWhole = convertToPartAndWhole(helper.TransOut, helper.MinorPerMajor)
	rec.IndusInPart, rec.IndusInWhole = convertToPartAndWhole(helper.IndusIn, helper.MinorPerMajor)
	rec.IndusOutPart, rec.IndusOutWhole = convertToPartAndWhole(helper.IndusOut, helper.MinorPerMajor)
	rec.RaseedbeforePart, rec.RaseedbeforeWhole = convertToPartAndWhole(helper.Raseedbefore, helper.MinorPerMajor)
	rec.RaseedPart, rec.RaseedWhole = convertToPartAndWhole(helper.Raseed, helper.MinorPerMajor)
}

func convertToPartAndWhole(orign float64, minor int) (float64, float64) {
	part := math.Mod(orign, float64(minor))
	whole := (orign - part) / float64(minor)

	return part, whole

}
