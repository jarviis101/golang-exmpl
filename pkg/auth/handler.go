package auth

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type loginRepsonse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var p Payload
	err := c.BindJSON(&p)
	if err != nil {
		fmt.Printf("login error: %s", err.Error())
	}

	token, err := getToken(&p)
	if err != nil {
		fmt.Printf("token error: %s", err.Error())
	}

	err = json.NewEncoder(c.Writer).Encode(loginRepsonse{
		Token: token,
	})
}
