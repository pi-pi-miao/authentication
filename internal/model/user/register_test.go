package user

import (
	"net/http"
	"runtime"
	"testing"
	"time"
)

var (
	name = "root"
	passwd = "123456"
	ip     = "192.168.0.100"
	port   = "3306"
	store  = "user"
	table  = "user"
)

func TestUser_Register(t *testing.T) {
	//u := &User{
	//	Userid:      "123",
	//	Name:        "pi-pi",
	//	Password:    "123456",
	//	Phone:       "18511396306",
	//	Email:       "pyrene@yeah.net",
	//	Gender:      "man",
	//	Age:         "18",
	//	Permission:  "0",
	//	List:        "0",
	//	MultiTenant: "0",
	//}
	//fmt.Println()
	//db.InitDb(name,passwd,ip,port,store)
	//err := u.Register()
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log("insert suc")
	//var now = "2021-1-20"
	//n,err := time.Parse("2006-01-02",now)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(n.Unix())
	//url := "http://www.baidu.com"
	//
	//for i:=0;i<6;i++ {
	//	//url := fmt.Sprintf("http://www.baidu.com ",i)
	//	// http.Get(url)
	//
	//	resp,err := http.Get(url)
	//	if err != nil {
	//		t.Log(err)
	//	}
	//	resp.Body.Close()
	//}
	//urlList := []string{"http://www.taobao.com", "http://www.baidu.com","http://www.cnblogs.com"}
	urlList1 := []string{"http://www.taobao.com"}
	for k,_ := range urlList1 {
		resp,err := http.Get(urlList1[k])
		if err != nil {
			t.Log(err)
		}
		resp.Body.Close()
	}
	time.Sleep(1*time.Second)
	t.Log(runtime.NumGoroutine())
	//time.Sleep(1*time.Second)
}
