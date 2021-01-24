package user

import (
	"authentication/pkg/config"
	"authentication/pkg/db"
	"authentication/pkg/sha"
	"authentication/pkg/uuid"
	"fmt"
	"log"
)

type User struct {
	Userid      string `json:"userid"`
	Name        string `json:"name" validate:"min=1,max=35"`
	Password    string `json:"password" validate:"required,min=1,max=50"`
	Phone       string `json:"phone" validate:"required,phone"`
	Email       string `json:"email" validate:"required,email"`
	Gender      string `json:"gender" validate:"required"`
	Age         string `json:"age" validate:"lte=200,gte=0"`
	Permission  string `json:"permission"`
	List        string `json:"list"`
	MultiTenant string `json:"tenant"`
}

func (u *User) Register() error {
	u.Userid = uuid.GetUuid()
	u.Permission = config.DomesticConsumer
	u.List = config.WhiteList
	u.Password = sha.SetSha256(u.Password)
	sql := fmt.Sprintf("insert into user (userid,name,password,phone,email,gender,age,permission,list,tenant) values('%v','%v','%v','%v','%v','%v','%v','%v','%v','%v')",
		u.Userid, u.Name, u.Password, u.Phone, u.Email, u.Gender, u.Age, u.Permission, u.List, u.MultiTenant)
	_, err := db.Insert(sql)
	if err != nil {
		log.Println("insert err", err)
		return err
	}
	return nil
}
