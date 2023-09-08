package service

import (
	"Immersive_dash/features/feedback"

	"github.com/go-playground/validator"
)

type feedbackService struct {
	feedbackData feedback.FeedbackDataInterface
	validate     *validator.Validate
}

// Create implements feedback.FeedbackServiceInterface.
func (service *feedbackService) Create(input feedback.Core) (feedback.Core, error) {
	// panic("unimplemented")
	result, err := service.feedbackData.Insert(input)
	if err != nil {
		return feedback.Core{}, err
	}
	return result, err
}

// Delete implements feedback.FeedbackServiceInterface.
func (service *feedbackService) Delete(id uint) error {
	panic("unimplemented")
}

// Update implements feedback.FeedbackServiceInterface.
func (service *feedbackService) Update(input feedback.Core) error {
	panic("unimplemented")
}

func New(repo feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &feedbackService{
		feedbackData: repo,
		validate:     validator.New(),
	}
}
