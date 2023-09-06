package mentee

import "time"

type Core struct {
	ID              uint
	Created_At      time.Time
	Updated_At      time.Time
	Deleted_At      time.Time
	FullName        string
	NickName        string
	Email           string
	Phone           int
	CurrentAddress  string
	HomeAddress     string
	Telegram        string
	ClassID         uint
	StatusID        uint
	Gender          string
	EducationType   string
	Major           string
	Graduate        string
	Institution     string
	EmergencyName   string
	EmergencyPhone  int
	EmergencyStatus string
	Status          StatusEntity
}

type StatusEntity struct {
	ID          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type MenteeDataInterface interface {
	Insert(input Core) error
	Delete(id uint) error
	SelectById(id uint) (Core, error)
}

type MenteeServiceInterface interface {
	Create(input Core) error
	Delete(id uint) error
	GetById(id uint) (Core, error)
}
