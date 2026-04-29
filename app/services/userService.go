package services

import (
	"edart-core/app/dtos"
	"edart-core/app/models"
	"edart-core/app/repositories"
	"time"
)

type UserService struct {
	repo * repositories.UserRespository
}

func NewUserService() *UserService {
	return &UserService{repo: repositories.NewUserRepo()}
}

func (a *UserService) CreateUserService(request *dtos.CreateUserRequest) (*models.UserInfo, error) {
	var userInfo models.UserInfo

	uid, err := a.repo.GetNextUserId()
	if err != nil {
		return nil, err
	}

	udob, err := time.Parse("2006-01-02", request.Dob)
	if err != nil {
		return nil, err
	}

	userInfo.Uid = uid
	userInfo.Fullname = request.Fullname
	userInfo.Country = request.Country
	userInfo.Dob = udob
	userInfo.IdNum = request.IdNum
	userInfo.IdType = request.IdType
	userInfo.Nationality = request.Nationality
	userInfo.CreateAt = time.Now()

	if err := a.repo.CreateUser(userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}