package handler

import (
	"Immersive_dash/app/helpers"
	"Immersive_dash/features/class"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService class.ClassServiceInterface
}

func NewClassHandler(service class.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) CreateClass(c echo.Context) error {
	name := c.FormValue("name")
	startDateStr := c.FormValue("start_date")
	graduateDateStr := c.FormValue("graduate_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid start date", nil))
	}

	graduateDate, err := time.Parse("2006-01-02", graduateDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid graduate date", nil))
	}

	core, err := handler.classService.CreateClass(name, startDate, graduateDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error creating class", err))
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "Success creating class", core))
}

func (handler *ClassHandler) GetClassByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid class ID", nil))
	}

	core, err := handler.classService.GetClassByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error getting class data", err))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success get class data", core))
}

func (handler *ClassHandler) UpdateClass(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid class ID", nil))
	}

	name := c.FormValue("name")
	startDateStr := c.FormValue("start_date")
	graduateDateStr := c.FormValue("graduate_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid start date", nil))
	}

	graduateDate, err := time.Parse("2006-01-02", graduateDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid graduate date", nil))
	}

	core, err := handler.classService.UpdateClass(uint(id), name, startDate, graduateDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error updating class data", err))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success update class data", core))
}

func (handler *ClassHandler) DeleteClassByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "Invalid class ID", nil))
	}

	err = handler.classService.DeleteClassByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "Error deleting class data", err))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "Success delete class data", nil))
}
