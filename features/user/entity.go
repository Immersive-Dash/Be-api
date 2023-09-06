package user

type Core struct {
	ID       uint
	FullName string
	TeamID   uint
	Team     TeamCore
	Password string
	Email    string
	Role     string
}

type TeamCore struct {
	ID   uint
	Name string
}

type UserDataInterface interface {
	Register(input Core) error
	Login(email string, password string) (dataLogin Core, err error)
	Read() ([]Core, error)
	Update(input Core) error
	ReadById(id uint) (Core, error)
	DeleteById(id uint) error
	UpdateById(id uint, input Core) error
}

type UserServiceInterface interface {
	CreateUser(input Core) error
	LoginUser(email string, password string) (dataLogin Core, token string, err error)
	GetUser() ([]Core, error)
	UpdateUser(input Core) error
	ReadUserById(id uint) (Core, error)
	DeleteUserById(id uint) error
	UpdateUserById(id uint, input Core) error
}
