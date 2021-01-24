package user

import (
	"authentication/pkg/db"
	"fmt"
	"log"
)

type LoginUser struct {
	Userid     string `json:"userid"`
	Phone      string `json:"phone" validate:"required,phone"`
	Password   string `json:"password" validate:"required,min=1,max=50"`
	Permission string `json:"permission"`
	List       string `json:"list"`
}

func (l *LoginUser) Login(phone string) error {
	sql := fmt.Sprintf("select * from %v where phone = %v", db.UserTable, phone)
	fmt.Println("sql is ==>", sql)
	if _, err := db.Get(sql, l); err != nil {
		log.Println("get user err", err)
		return err
	}
	return nil
}
