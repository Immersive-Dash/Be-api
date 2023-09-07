package feedback

import _dataMentee "Immersive_dash/features/mentee"

type Core struct {
	ID       uint
	Notes    string
	MenteeID uint
	Mentee   _dataMentee.Core
	Status   string
}

type FeedbackDataInterface interface {
	Insert(input Core) (Core, error)
	Update(input Core) error
	Delete(id uint) error
}

type FeedbackServiceInterface interface {
	Create(input Core) (Core, error)
	Update(input Core) error
	Delete(id uint) error
}
