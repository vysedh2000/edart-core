package dtos

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Nationality string `json:"nationality"`
	IdNum string `json:"idnum"`
	IdType string `json:"idtype"`
	Country string `json:"country"`
	Dob string `json:"dob"`
}