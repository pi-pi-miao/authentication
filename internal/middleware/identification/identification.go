package identification

import (
	"authentication/pkg/base64"
	"authentication/pkg/cache"
	"authentication/pkg/config"
	"authentication/pkg/sha"
	"authentication/pkg/str"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)



type Auth struct {
	Uid  string    `json:"userid"`
}

func Identification(c *gin.Context){
	baseToken := c.Request.Header.Get("token")
	if baseToken == "" {
		fmt.Println("[Identification] ---")
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	// is Blacklist ?
	if cache.Cache.Get(baseToken) {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	byteToken,err  := base64.BaseDecode(baseToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	token := str.BytesToString(byteToken)
	if token[len(token)-2:len(token)-1] == config.BlackList {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}

	t,err := strconv.ParseInt(token[len(token)-12:len(token)-2],10,64)
	if err != nil {
		log.Println("[Identification] string to int64 err",err)
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	if time.Now().Unix() -t  > 300 {
		c.Redirect(http.StatusOK,"127.0.0.1:8090/login")
		c.Abort()
		return
	}

	body,err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("[Identification] read request body err",err)
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}
	user := &Auth{}
	if err := json.Unmarshal(body,user);err != nil {
		log.Println("[Identification] unmarshal request body",err)
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		c.Abort()
		return
	}

	sha.GetSha256(fmt.Sprintf("%v%v",config.Salt,user.Uid),token[len(token)-2:len(token)-1],token[len(token)-1:len(token)],token[len(token)-12:len(token)-2])


	c.Next()
}
