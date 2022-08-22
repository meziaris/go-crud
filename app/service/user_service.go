package service

import (
	"errors"
	"go-crud/app/model/domain"
	"go-crud/app/model/request"
	"go-crud/app/model/response"
	"go-crud/app/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
	authService    *authService
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: &userRepository}
}

func (service *UserService) Register(request request.RegisterRequest) (userResponse response.UserResponse, err error) {
	user := domain.User{
		Name:      request.Name,
		Username:  request.Username,
		Email:     request.Email,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userRes, _ := service.UserRepository.FindByUsername(request.Username)
	if userRes.Username != "" {
		return userResponse, errors.New("username has been taking")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return userResponse, err
	}
	user.Password = string(password)

	rsp, err := service.UserRepository.Insert(user)
	if err != nil {
		return userResponse, err
	}

	userResponse = response.UserResponse{
		Id:       rsp.Id,
		Name:     rsp.Name,
		Email:    rsp.Email,
		Username: rsp.Username,
		Role:     rsp.Role,
	}

	return userResponse, nil
}

func (service *UserService) Login(request request.LoginRequest) (loginResponse response.LoginResponse, err error) {
	username := request.Username
	password := request.Password

	user, err := service.UserRepository.FindByUsername(username)
	if err != nil {
		return loginResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return loginResponse, errors.New("incorrect password")
	}

	token, err := service.authService.JwtGenerateToken(user.Id)
	if err != nil {
		return loginResponse, err
	}

	loginResponse = response.LoginResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}

	return loginResponse, nil
}

func (service *UserService) ListAll(id primitive.ObjectID) (responses []response.ShowAll, err error) {
	role, _ := service.UserRepository.FindById(id)
	if role.Role != "admin" {
		return responses, errors.New("you need permission to perform this action")
	}

	users, err := service.UserRepository.FindAll()
	if err != nil {
		return responses, err
	}

	for _, user := range users {
		responses = append(responses, response.ShowAll{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
		})
	}

	return responses, nil
}

func (service *UserService) UserDetail(id primitive.ObjectID) (userResponse response.UserResponse, err error) {
	user, err := service.UserRepository.FindById(id)
	if err != nil {
		return userResponse, err
	}

	userResponse = response.UserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
	}

	return userResponse, nil
}

func (service *UserService) Update(id primitive.ObjectID, username string, request request.UpdateRequest) (userResponse response.UserResponse, err error) {

	user, err := service.UserRepository.FindByUsername(username)
	if err != nil {
		return userResponse, err
	}

	role, _ := service.UserRepository.FindById(id)
	if role.Role != "admin" {
		return userResponse, errors.New("you need permission to perform this action")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.UpdatedAt = time.Now()

	rsp, err := service.UserRepository.Update(user)
	if err != nil {
		return userResponse, err
	}

	userResponse = response.UserResponse{
		Id:       rsp.Id,
		Name:     rsp.Name,
		Email:    rsp.Email,
		Username: rsp.Username,
		Role:     rsp.Role,
	}

	return userResponse, nil
}

func (service *UserService) Delete(id primitive.ObjectID, username string) (message string, err error) {
	role, _ := service.UserRepository.FindById(id)
	if role.Role != "admin" {
		return "opps, somethting when wrong!", errors.New("you need permission to perform this action")
	}

	user, _ := service.UserRepository.FindByUsername(username)

	err = service.UserRepository.Delete(user.Id)
	if err != nil {
		return "opps, somethting when wrong!", err
	}

	return "delete user success!", nil
}
