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
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid request", nil))
	}

	coreMentee := RequestToCore(*input)
	err := handler.menteeService.Create(coreMentee)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error insert data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "Success insert data", nil))
}

func (handler *MenteeHandler) DeleteMenteeByID(c echo.Context) error {

	id := c.Param("id_mentee")

	Userid := middlewares.ExtractTokenRole(c)
	if Userid != "admin" {
		return c.JSON(http.StatusForbidden, "forbiden access")
	}
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, "wrong id")
	}

	err := handler.menteeService.Delete(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error")
	}

	return c.JSON(http.StatusOK, "succes")
}

func (handler *MenteeHandler) GetMenteeByID(c echo.Context) error {

	idMenteeStr := c.Param("id_mentee")
	idMentee, errMentee := strconv.Atoi(idMenteeStr)
	if errMentee != nil {
		return c.JSON(http.StatusBadRequest, "Invalid mentee ID")
	}

	Userid := middlewares.ExtractTokenRole(c)
	if Userid != "admin" {
		return c.JSON(http.StatusForbidden, "Forbidden access")
	}

	result, err := handler.menteeService.GetById(uint(idMentee))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error getting mentee data")
	}

	resultResponse := MenteeResponse{
		ID:       result.ID,
		FullName: result.FullName,
		Email:    result.Email,
		Phone:    result.Phone,
		Telegram: result.Telegram,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get mentee data", resultResponse))
}

func (handler *MenteeHandler) UpdateMentee(c echo.Context) error {
	// Mengambil ID mentee dari URL parameter
	idMenteeStr := c.Param("id_mentee")
	idMentee, errMentee := strconv.Atoi(idMenteeStr)
	if errMentee != nil {
		return c.JSON(http.StatusBadRequest, "Invalid mentee ID")
	}

	Userid := middlewares.ExtractTokenRole(c)
	if Userid != "admin" {
		return c.JSON(http.StatusForbidden, "Forbidden access")
	}

	input := new(MenteeRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid request", nil))
	}

	coreMentee := RequestToCore(*input)

	err := handler.menteeService.Update(uint(idMentee), coreMentee)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating mentee data", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update mentee data", nil))
}

func (handler *MenteeHandler) ReadMentee(c echo.Context) error {
	result, err := handler.menteeService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var menteeResponse []MenteeResponse
	for _, value := range result {
		menteeResponse = append(menteeResponse, MenteeResponse{
			ID:            value.ID,
			Class:         value.Class,
			FullName:      value.FullName,
			NickName:      value.NickName,
			Status:        value.Status,
			EducationType: value.EducationType,
			Gender:        value.Gender,
			Email:         value.Email,
			Phone:         value.Phone,
			Telegram:      value.Telegram,
		})
		// fmt.Println("data: ", userResponse)
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", menteeResponse))
}
