package user

import (
	"authentication/internal/model/user"
	"authentication/pkg/config"
	"authentication/pkg/sha"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type response struct {

}

func Login(c *gin.Context){
	fmt.Println("[login]")
	body,err:= ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	u := &user.LoginUser{}
	if err := json.Unmarshal(body,u);err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}

	user := &user.LoginUser{}
	if err := user.Login(u.Phone);err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	if sha.SetSha256(u.Password) != user.Password {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	if user.List == config.BlackList {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	token := sha.GetSha256(fmt.Sprintf("%v%v",config.Salt,user.Userid),user.List,user.Permission,strconv.FormatInt(time.Now().Unix(),10))
	c.Writer.Header().Set("token",token)
	c.JSON(http.StatusOK,gin.H{"status":"succ"})
	return
}
