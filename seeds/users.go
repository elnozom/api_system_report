package seeds

import (
	"github.com/ahmedashrafdev/reports/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *gorm.DB) error {
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	password2, _ := bcrypt.GenerateFromPassword([]byte("Ihf22/10/2021"), bcrypt.DefaultCost)

	db.Create(&model.User{Email: "local-ghonim@elnozom.com", ServerID: 1, Password: string(password)})
	db.Create(&model.User{Email: "halal@elnozom.com", ServerID: 2, Password: string(password)})
	db.Create(&model.User{Email: "dental@elnozom.com", ServerID: 3, Password: string(password)})
	db.Create(&model.User{Email: "local@elnozom.com", ServerID: 4, Password: string(password)})
	db.Create(&model.User{Email: "ghonim@elnozom.com", ServerID: 5, Password: string(password)})
	db.Create(&model.User{Email: "hanin@elnozom.com", ServerID: 6, Password: string(password2)})
	return db.Create(&model.User{Email: "rady@elnozom.com", ServerID: 7, Password: string(password)}).Error
}
