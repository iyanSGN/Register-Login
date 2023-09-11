package handlers

import (
	"fmt"
	"goLANG/app"
	"goLANG/helpers"
	"goLANG/models"
	"goLANG/token"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	user := models.MasterUser{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := helpers.HashPassword(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	createdUser, err := app.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userID := uint(user.Id)
	token, err := token.GenerateToken(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"status_code" : http.StatusOK,
		"token"	:	token,
		"data" : map[string]interface{}{
		"message": "register successfull",
		"id":	createdUser.Id,
		"name":	createdUser.Name,
		"email": createdUser.Email,
		"password": createdUser.Password,
		},
	}

	return c.JSON(http.StatusOK, response)

}

func GetUser(c echo.Context) error {
	GetUser, err := app.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data" : GetUser,
		"status_code" : http.StatusOK,
	})


}

func UpdateUser(c echo.Context) error {
    idUser := c.Param("id")
    id, err := strconv.Atoi(idUser)
    if err != nil {
        return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
    }

    // Mendapatkan data pengguna yang akan diperbarui dari permintaan
    updatedUser := models.MasterUser{}
    if err := c.Bind(&updatedUser); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    // Memeriksa apakah ada password baru yang akan dienkripsi
    if updatedUser.Password != "" {
        // Enkripsi password yang baru
        if err := helpers.HashPassword(&updatedUser); err != nil {
            return c.JSON(http.StatusInternalServerError, err.Error())
        }
    }

    err = app.UpdateUser(id, updatedUser)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "message":     "User Updated Successfully",
        "status_code": http.StatusOK,
	})
}

func DeleteUser(c echo.Context) error {
	idUser := c.Param("id")

	id, err := strconv.Atoi(idUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("invalid id"))
	}

	err = app.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "User id has been deleted",
		"status_code" : http.StatusOK,
	})
}


