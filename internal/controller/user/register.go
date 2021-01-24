package user

import (
	user2 "authentication/internal/model/user"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
)

func Register(c *gin.Context){
	fmt.Println("[register]")
	body,err:= ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	user := &user2.User{}
	if err := json.Unmarshal(body,user);err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	validate := validator.New()
	if err := validate.Struct(user);err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	if err := user.Register();err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"status":"suc"})
	return
}
