package handler

import (
	"Immersive_dash/app/helpers"
	"Immersive_dash/app/middlewares"
	"Immersive_dash/features/mentee"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	menteeService mentee.MenteeServiceInterface
}

func New(service mentee.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: service,
	}
}

func (handler *MenteeHandler) CreateMentee(c echo.Context) error {

	input := new(MenteeRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	coreMentee := RequestToCore(*input)
	err := handler.menteeService.Create(coreMentee)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {

			return c.JSON(http.StatusBadRequest, nil)
		} else {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}

	return c.JSON(http.StatusCreated, "success insert data")
}

func (handler *MenteeHandler) DeleteMenteeByID(c echo.Context) error {

	Userid := middlewares.ExtractTokenRole(c)
	if Userid != "admin" {
		return c.JSON(http.StatusForbidden, "forbiden access")
	}
	idConv, errConv := strconv.Atoi(Userid)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "wrong id")
	}

	err := handler.menteeService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}

	return c.JSON(http.StatusOK, "succes")
}

func (handler *MenteeHandler) GetMenteeById(c echo.Context) error {
	Userid := middlewares.ExtractTokenRole(c)
	if Userid != "admin" {
		return c.JSON(http.StatusForbidden, "forbiden access")
	}
	idConv, errConv := strconv.Atoi(Userid)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "wrong id")
	}

	err := handler.menteeService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}

	result, err := handler.menteeService.GetById(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, "Error")
		} else {
			return c.JSON(http.StatusInternalServerError, "error insert data")

		}
	}
	resultResponse := MenteeResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Email:    result.Email,
		Phone:    result.Phone,
		Telegram: result.Telegram,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", resultResponse))
}
