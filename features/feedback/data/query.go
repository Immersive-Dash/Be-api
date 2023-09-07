package data

import (
	"Immersive_dash/features/feedback"

	"gorm.io/gorm"
)

type feedbackQuery struct {
	db *gorm.DB
}

// Delete implements feedback.FeedbackDataInterface.
func (repo *feedbackQuery) Delete(id uint) error {
	panic("unimplemented")
}

// Insert implements feedback.FeedbackDataInterface.
func (repo *feedbackQuery) Insert(input feedback.Core) (feedback.Core, error) {
	feedbackData := CoreToModel(input)
	tx := repo.db.Create(&feedbackData)
	if tx.Error != nil {
		return feedback.Core{}, tx.Error
	}
	return ModelToCore(feedbackData), nil
}

// Update implements feedback.FeedbackDataInterface.
func (repo *feedbackQuery) Update(input feedback.Core) error {
	panic("unimplemented")
}

func New(db *gorm.DB) feedback.FeedbackDataInterface {
	return &feedbackQuery{
		db: db,
	}
}
