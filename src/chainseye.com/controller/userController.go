package controller

import (
	"net/http"
	"time"
    "math/rand"

	"github.com/gin-gonic/gin"
)

//GetUserInfoByUID 根据UID获取基本信息
func GetUserInfoByUID(context *gin.Context) {

    str := "0123456789abcdef"  
    bytes := []byte(str)  
    result := []byte{}  
    r := rand.New(rand.NewSource(time.Now().UnixNano()))  
    for i := 0; i < 6; i++ {  
        result = append(result, bytes[r.Intn(len(bytes))])  
    }  
	// ur, err := service.GetUserInfoByUID(1)
	// if err != nil {
	// 	context.JSON(http.StatusOK, gin.H{
	// 		"retCode": 1001,
	// 		"info":    err,
	// 	})
	// 	return
	// }
	// userData := map[string]interface{}{
	// 	"mobile":   ur.Mobile,
	// 	"password": ur.Password,
	// }
	context.JSON(http.StatusOK, gin.H{
		"retCode": 0,
		"token": string(result),
		// "info":    userData,
	})
}
