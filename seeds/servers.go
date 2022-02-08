package seeds

import (
	"github.com/ahmedashrafdev/golang-echo-realworld-example-app/model"
	"github.com/jinzhu/gorm"
)

func CreateServer(db *gorm.DB) error {
	db.Create(&model.Server{
		DbUser:     "pda",
		DbPassword: "123",
		DbIP:       "nozom.myfirewall.co",
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
		DbIP:       "nozom.myfirewall.co",
		DbName:     "mcstree",
		ServerName: "local",
	})
	db.Create(&model.Server{
		DbUser:     "pda",
		DbPassword: "123",
		DbIP:       "ghoneimmarket.myfirewall.co",
		DbName:     "Ghonim06",
		ServerName: "ghonim",
	})
	db.Create(&model.Server{
		DbUser:     "mynozom",
		DbPassword: "Ceva@2290100",
		DbIP:       "merwesh.myfirewall.co",
		DbName:     "STOCK",
		ServerName: "merwesh",
	})
	return db.Create(&model.Server{
		DbUser:     "SqlAuth",
		DbPassword: "lz6RGWdFnI*uB7@xXdQ",
		DbIP:       "41.41.177.180",
		DbName:     "STOCK2021",
		ServerName: "elrady",
	}).Error
}
