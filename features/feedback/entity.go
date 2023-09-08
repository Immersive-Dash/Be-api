package feedback

type Core struct {
	ID       uint
	Notes    string
	MenteeID uint
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
