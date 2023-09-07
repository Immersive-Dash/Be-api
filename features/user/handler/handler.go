package handler

import (
	"Immersive_dash/app/helpers"
	"Immersive_dash/app/middlewares"
	"Immersive_dash/features/user"
	"net/http"
	"strconv"
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
			Role:     value.Role,
		})
		// fmt.Println("data: ", userResponse)
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

func (handler *userHandler) RegisterUser(c echo.Context) error {
	userInput := new(UserRequest)
	errBind := c.Bind(&userInput)
	role := middlewares.ExtractTokenRole(c)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "access denied", nil))
	}
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	//mapping dari struct request to struct core
	userCore := RequestToCore(*userInput)
	result, err := handler.userService.CreateUser(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	registerResponse := UserRegisterResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Role:     result.Role,
		Email:    result.Email,
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", registerResponse))
}

func (handler *userHandler) UpdateUser(c echo.Context) error { // update user yang login
	role := middlewares.ExtractTokenRole(c)
	userID := middlewares.ExtractTokenUserId(c)
	userInput := new(UserUpdateRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	userCore := UpdateRequestToCore(*userInput)
	result, err := handler.userService.UpdateUser(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}
	updateResponse := UserResponse{
		ID:       uint(userID),
		FullName: result.FullName,
		Team:     result.Team.Name,
		Email:    result.Email,
		Role:     role,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update data", updateResponse))
}

func (handler *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id_user")
	role := middlewares.ExtractTokenRole(c)
	if role != "admin" {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "access denied", nil))
	}
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error id not valid", nil))
	}
	err := handler.userService.DeleteUserById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete data", nil))

}

func (handler *userHandler) GetUserById(c echo.Context) error {
	id := c.Param("id_user")
	role := middlewares.ExtractTokenRole(c)
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error id not valid", nil))
	}

	result, err := handler.userService.ReadUserById(uint(idConv))
	if err != nil {

		return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "data not found", nil))
	}
	// mapping dari struct core to struct response
	resultResponse := UserResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Team:     result.Team.Name,
		Email:    result.Email,
		Role:     role,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", resultResponse))
}

func (handler *userHandler) UpdateById(c echo.Context) error {
	userInput := new(UserRequest)
	role := middlewares.ExtractTokenRole(c)
	id := c.Param("id_user")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error data id. data not valid", nil))
	}
	errBind := c.Bind(&userInput) // mendapatkan data yang dikirim oleh FE melalui request body
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	userCore := RequestToCore(*userInput)
	result, err := handler.userService.UpdateUserById(uint(idParam), userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	updateResponse := UserResponse{
		ID:       uint(idParam),
		FullName: result.FullName,
		Team:     result.Team.Name,
		Email:    result.Email,
		Role:     role,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success update data", updateResponse))
}
