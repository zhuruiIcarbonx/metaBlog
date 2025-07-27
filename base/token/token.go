package token

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetUserId(c *gin.Context) uint {

	anyValue, _ := c.Get("userId")
	log.Println("anyValue------------", anyValue)
	var userId int
	switch v := anyValue.(type) {
	case int:
		userId = v
	case float64:
		userId = int(v)
	default:
		log.Printf("userId type is not supported: %T", v)
		return 0
	}
	uId := uint(userId)
	return uId

}

func GetUsername(c *gin.Context) string {

	anyValue, _ := c.Get("username")
	log.Println("anyValue------------", anyValue)
	var username string
	switch v := anyValue.(type) {
	case string:
		username = v
	default:
		log.Printf("userId type is not supported: %T", v)
		username = ""
	}
	return username

}
