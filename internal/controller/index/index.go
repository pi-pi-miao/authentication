package index

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Index(c *gin.Context){
	log.Println("suc")
	c.JSON(http.StatusOK,gin.H{"status":"suc"})
	return
}
