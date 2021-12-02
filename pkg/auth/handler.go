package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var p Payload
	jwtService := JWTAuthService()
	err := c.BindJSON(&p)
	if err != nil {
		fmt.Printf("login error: %s", err.Error())
	}

	token, err := jwtService.Get(&p)
	if err != nil {
		fmt.Printf("token error: %s", err.Error())
	}

	c.JSON(http.StatusOK, &loginResponse{
		Token: token,
	})
}

func Logout(c *gin.Context) {

}
