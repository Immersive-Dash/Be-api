package handler

import (
	"Immersive_dash/app/helpers"
	"Immersive_dash/features/user"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *userHandler {
	return &userHandler{
		userService: service,
	}
}

func (handler *userHandler) ReadUser(c echo.Context) error {
	result, err := handler.userService.GetUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var userResponse []UserResponse
	for _, value := range result {
		userResponse = append(userResponse, UserResponse{
			ID:       value.ID,
			FullName: value.FullName,
			Team:     value.Team.Name,
			Email:    value.Email,
		})
		fmt.Println("data: ", userResponse)
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", userResponse))
}

func (handler *userHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	dataLogin, token, err := handler.userService.LoginUser(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	response := map[string]any{
		"token": token,
		"role":  dataLogin.Role,
		"id":    dataLogin.ID,
		"email": dataLogin.Email,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

// func (handler *userHandler) Register(c echo.Context) error {
// 	userInput := new(user.Core)
// 	errBind := c.Bind(&userInput)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
// 	}
// }
