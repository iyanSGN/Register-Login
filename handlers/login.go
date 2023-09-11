package handlers

import (
	"fmt"
	"goLANG/app"
	"goLANG/token"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func LoginAccount(c echo.Context) error {
	loginData := struct {
		Email	string 	`json:"email"`
		Password string `json:"password"`
	}{}
	
	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid request payload"))
	}

	user, err := app.GetUserByEmail(loginData.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("invalid email/username, try again"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userID := uint(user.Id)
	token, err := token.GenerateToken(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to generate token"))
	}

	if c.Get("newToken") != nil {
		newToken, ok := c.Get("newToken").(string)
		if ok && newToken != token {
			token = newToken
		}
	}

	response := map[string]interface{}{
		"Status_code" : http.StatusOK,
		"token": token,
		"data" : map[string]interface{}{
			"message" : "Login Successfull",
			"Id" : 	user.Id,
			"Email" : user.Email,
		},
	}

	return c.JSON(http.StatusOK, response)
}