package feedback

type Core struct {
	ID     uint   `json:"id"`
	Notes  string `json:"notes"`
	Status string `json:"status"`
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
