package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var db  *xorm.Engine
var TestDb *xorm.Engine

func InitDb(name,passwd,ip,port,store string)error{
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",name,passwd,ip,port,store))
	if err != nil {
		fmt.Println("[initDb] init db err",err)
		return err
	}
	engine.SetMaxIdleConns(50)
	engine.SetMaxOpenConns(100)
	if err := engine.Ping();err != nil {
		fmt.Println("[initDb] ping db err",err)
		return err
	}
	db = engine
	TestDb = engine
	return nil
}


func Insert(data string)(int64,error){
	res,err := db.Exec(data)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}


func Get(sql string,data interface{})(bool,error){
	return db.SQL(sql).Get(data)
}
