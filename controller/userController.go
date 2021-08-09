package controller

import (
	"net/http"
	"restfulAPI/lib/database"
	"restfulAPI/middlewares"
	"restfulAPI/models"
	"strconv"

	"github.com/labstack/echo"
)

/////////////////USER/////////////////

func SignUpUser(c echo.Context) error {
	addUser := models.User{}
	c.Bind(&addUser)
	user, err := database.CreateUser(addUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot insert data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "new user added",
		"data":    user,
	})
}

func Login(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)

	user, err := database.LoginUsers(userData.Email, userData.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "username or password is not correct",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Welcome",
		"users":   user,
	})
}

func ChangeProfile(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	editUser, err := database.GetOneUser(id)
	loggedInUserId, token := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != id || editUser.Token != token {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}

	c.Bind(&editUser)
	user, err := database.EditUser(editUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot edit data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile Updated!",
		"data":    user,
	})
}

func Logout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	logout, err := database.GetOneUser(id)
	logout.Token = ""
	c.Bind(&logout)
	user, err := database.EditUser(logout)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot logout",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "GOODBYE!",
		"data":    user,
	})
}

func ShowProfile(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	userAuth, err := database.GetOneUser(userId)
	loggedInUserId, token := middlewares.ExtractTokenUserId(c)
	if loggedInUserId != userId || userAuth.Token != token {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	user, err := database.GetDetailUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find the user",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}
