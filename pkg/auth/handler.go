package auth

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessfulResponseData struct {
	Token string `json:"token"`
}

func Authenticate(c *gin.Context) {
	w := c.Writer
	nickname := "lex"

	token, err := getToken(nickname)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorResponseData{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		})
		w.Header().Add("Error", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&SuccessfulResponseData{
		Token: token,
	})
}
