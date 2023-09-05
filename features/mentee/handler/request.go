package api

type CreateMenteeRequest struct {
	IDClass         int    `json:"id_class"`
	FullName        string `json:"full_name"`
	NickName        string `json:"nick_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	CurrentAddress  string `json:"current_address"`
	HomeAddress     string `json:"home_address"`
	Telegram        string `json:"telegram"`
	Gender          string `json:"gender"`
	EducationType   string `json:"education_type"`
	Major           string `json:"major"`
	Graduate        string `json:"graduate"`
	Institution     string `json:"institution"`
	EmergencyName   string `json:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status"`
	Status          string `json:"status"`
}

type UpdateMenteeRequest struct {
	IDClass         int    `json:"id_class"`
	FullName        string `json:"full_name"`
	NickName        string `json:"nick_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	CurrentAddress  string `json:"current_address"`
	HomeAddress     string `json:"home_address"`
	Telegram        string `json:"telegram"`
	Gender          string `json:"gender"`
	EducationType   string `json:"education_type"`
	Major           string `json:"major"`
	Graduate        string `json:"graduate"`
	Institution     string `json:"institution"`
	EmergencyName   string `json:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status"`
	Status          string `json:"status"`
}
