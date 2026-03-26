package dtos

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Nationality string `json:"nationality"`
	Country string `json:"country"`
	Dob string `json:"dob"`
}