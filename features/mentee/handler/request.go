
package handler

import "Immersive_dash/features/mentee"

type MenteeRequest struct {
	StatusID        uint   `json:"status_id,omitempty" form:"status_id"`
	ClassID         uint   `json:"class_id,omitempty" form:"class_id"`
	FullName        string `json:"full_name,omitempty" form:"full_name"`
	NickName        string `json:"nick_name,omitempty" form:"nick_name"`
	CurrentAddress  string `json:"current_address,omitempty" form:"current_address"`
	HomeAddress     string `json:"home_address,omitempty" form:"home_address"`
	Email           string `json:"email,omitempty" form:"email"`
	Gender          string `json:"gender,omitempty" form:"gender"`
	Telegram        string `json:"telegram,omitempty" form:"telegram"`
	Phone           string `json:"phone,omitempty" form:"phone"`
	EmergencyName   string `json:"emergency_name,omitempty" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone,omitempty" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status,omitempty" form:"emergency_status"`
	EducationType   string `json:"education_type,omitempty" form:"education_type"`
	Major           string `json:"major,omitempty" form:"major"`
	Institution     string `json:"institution,omitempty" form:"institution"`
	Graduate        string `json:"graduate,omitempty" form:"graduate"`
}

func RequestToCore(input MenteeRequest) mentee.Core {
	return mentee.Core{
		StatusID:        input.StatusID,
		ClassID:         input.ClassID,
		FullName:        input.FullName,
		NickName:        input.NickName,
		CurrentAddress:  input.CurrentAddress,
		HomeAddress:     input.HomeAddress,
		Email:           input.Email,
		Gender:          input.Gender,
		Telegram:        input.Telegram,
		EmergencyName:   input.EmergencyName,
		EmergencyStatus: input.EmergencyStatus,
		EducationType:   input.EducationType,
		Major:           input.Major,
		Institution:     input.Institution,
		Graduate:        input.Graduate,
	}

