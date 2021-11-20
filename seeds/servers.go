package seeds

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
)

func CreateServer(db *gorm.DB) error {
	db.Create(&model.Server{
		DbUser:     "pda",
		DbPassword: "123",
		DbIP:       "192.168.1.167",
		DbName:     "Ghonim06",
		ServerName: "local-ghonim",
	})

	db.Create(&model.Server{
		DbUser:     "mcs",
		DbPassword: "123",
		DbIP:       "41.38.87.59",
		DbName:     "stock_main",
		ServerName: "halal",
	})
	db.Create(&model.Server{
		DbUser:     "web",
		DbPassword: "123",
		DbIP:       "dental.myfirewall.co",
		DbName:     "STOCK2021",
		ServerName: "dental",
	})
	db.Create(&model.Server{
		DbUser:     "mcs",
		DbPassword: "123",
		DbIP:       "192.168.1.167",
		DbName:     "mcstree",
		ServerName: "local",
	})
	db.Create(&model.Server{
		DbUser:     "pda",
		DbPassword: "123",
		DbIP:       "nozom.myfirewall.co",
		DbName:     "Ghonim06",
		ServerName: "ghonim",
	})
	return db.Create(&model.Server{
		DbUser:     "SqlAuth",
		DbPassword: "lz6RGWdFnI*uB7@xXdQ",
		DbIP:       "41.41.177.180",
		DbName:     "STOCK2021",
		ServerName: "elrady",
	}).Error
}
