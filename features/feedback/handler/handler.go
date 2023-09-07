package handler

import (
	"Immersive_dash/app/helpers"
	"Immersive_dash/features/feedback"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	feedbackService feedback.FeedbackServiceInterface
}

func New(service feedback.FeedbackServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		feedbackService: service,
	}
}

func (handler *FeedbackHandler) CreateFeedback(c echo.Context) error {
	userInput := new(FeedbackRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	feedbackCore := RequestToCore(*userInput)
	result, err := handler.feedbackService.Create(feedbackCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	feedbackResponse := feedbackResponse{
		ID:     result.ID,
		Status: result.Status,
		Notes:  result.Notes,
	}
	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", feedbackResponse))
}
