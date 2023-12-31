package mentee

import (
	"Immersive_dash/features/feedback"
	"time"
)

type Core struct {
	ID             uint
	Created_At     time.Time
	Updated_At     time.Time
	Deleted_At     time.Time
	FullName       string
	NickName       string
	Email          string
	Phone          string
	CurrentAddress string
	HomeAddress    string
	Telegram       string
	ClassID        uint
	Class          string
	//StatusID        uint
	Gender          string
	EducationType   string
	Major           string
	Graduate        string
	Institution     string
	EmergencyName   string
	EmergencyPhone  string
	EmergencyStatus string
	Status          string
	Feedbacks       []feedback.Core
}

// type ClassCore struct {
// 	ID   uint
// 	Name string
// }

// type StatusEntity struct {
// 	ID        uint
// 	Name      string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt time.Time
// }

type MenteeDataInterface interface {
	Insert(input Core) error
	Delete(id uint) error
	SelectById(id uint) (Core, error)
	Update(data Core) error
	Select(class string, status string, category string) ([]Core, error)
	SelectMenteeFeedback(id uint) ([]Core, error)
}

type MenteeServiceInterface interface {
	Create(input Core) error
	Delete(id uint) error
	GetById(id uint) (Core, error)
	Update(id uint, newData Core) error
	GetAll(class string, status string, category string) ([]Core, error)
	GetMenteeFeedback(id uint) ([]Core, error)
}
