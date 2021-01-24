package cache

import (
	"sync"
	"time"
)

var (
	Cache *cache
	once = &sync.Once{}
)


type cache struct {
	Data map[string]int64
	Lock *sync.RWMutex
}

func newCache(){
	once.Do(func() {
		Cache = &cache{
			Data: make(map[string]int64,1024),
			Lock: &sync.RWMutex{},
		}
		go daemon()
	})
}

func daemon(){
	Cache.Lock.RLock()
	data := Cache.Data
	Cache.Lock.RUnlock()
	t := time.Now().Unix()
	for k,_ := range data {
		if t - data[k] >= 300 {
			Cache.Delete(k)
		}
	}
	time.Sleep(1*time.Minute)
}

func (c *cache)Get(token string)bool {
	newCache()
	c.Lock.RLock()
	if _,ok := c.Data[token];ok {
		c.Lock.RUnlock()
		return true
	}
	c.Lock.RUnlock()
	return false
}

func (c *cache)Set(token string,t int64){
	newCache()
	c.Lock.Lock()
	c.Data[token] = t
	c.Lock.Unlock()
}


func (c *cache)Delete(k string){
	newCache()
	Cache.Lock.Lock()
	delete(Cache.Data, k)
	Cache.Lock.Unlock()
	m := make(map[string]int64,len(Cache.Data))
	Cache.Lock.Lock()
	Cache.Data = m
	Cache.Lock.Unlock()
}
