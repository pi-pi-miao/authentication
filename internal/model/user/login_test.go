package user

import (
	"log"
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//var intMap map[int]int
//var cnt = 8192

var lastTotalFreed uint64
var intMap sync.Map
var cnt = 80920

func TestLoginUser_Login(t *testing.T) {
	//l := &LoginUser{
	//	Phone:      "18511396306",
	//}
	//db.InitDb(name,passwd,ip,port,store)
	//if err := l.Login(l.Phone);err != nil {
	//	t.Log(err)
	//	return
	//}
	//s := "hello"
	//fmt.Println([]rune(s))
	//fmt.Println(l)
	//text := `世界！123 .`

	// 查找连续的小写字母
	//reg := regexp.MustCompile(`[a-z]+`)
	////fmt.Println(reg.FindAllString(text,-1))
	//t.Log(reg.FindString(text))
	//t.Log(reg.MatchString(text))

	//s := []string{}
	//if len(s) == 0 || s[0] == "" {
	//	t.Log("ok")
	//}
	//t.Log(100|2)
	//s := []string{"1","2","3"}
	//t.Log(strings.Join(s,"."))
	//srcPath := "/home/pi-pi/project/template/example/hehe"
	//destPath := "/home/pi-pi/project/template/example/hello"
	//if err := os.Rename(srcPath,destPath);err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log("ok")
	do1()
}

//func initMap(){
//	intMap = make(map[int]int, cnt)
//
//	for i := 0; i < cnt; i++ {
//		intMap[i] = i
//	}
//}

func initMap1(){

	for i := 0; i < cnt; i++ {
		intMap.Store(i, i)
	}
}


func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v  Just Freed = %v Sys = %v NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc}

//func do(){
//	printMemStats()
//
//	initMap()
//	runtime.GC()
//	printMemStats()
//
//	log.Println(len(intMap))
//	for i := 0; i < cnt; i++ {
//		delete(intMap, i)
//	}
//	log.Println(len(intMap))
//
//	runtime.GC()
//	printMemStats()
//
//	intMap = nil
//	runtime.GC()
//	printMemStats()
//}

func do1(){
	printMemStats()

	initMap1()
	runtime.GC()
	printMemStats()

	//fmt.Printf("%#v", intMap)
	fmt.Println()

	for i := 0; i < cnt; i++ {
		intMap.Delete(i)
	}
	//fmt.Printf("%#v", intMap)
	fmt.Println()

	runtime.GC()
	printMemStats()

}