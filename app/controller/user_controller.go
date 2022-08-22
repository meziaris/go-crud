package controller

import (
	"encoding/json"
	"go-crud/app/helper"
	"go-crud/app/model/request"
	"go-crud/app/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userController struct {
	UserService *service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{UserService: &service}
}

func (controller *userController) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := request.RegisterRequest{}
	err := decoder.Decode(&user)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	userResponse, err := controller.UserService.Register(user)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	helper.WebResponse(w, http.StatusOK, "OK", userResponse)
}

func (controller *userController) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	loginRequest := request.LoginRequest{}
	err := decoder.Decode(&loginRequest)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err)
		return
	}

	loginResponse, err := controller.UserService.Login(loginRequest)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	helper.WebResponse(w, http.StatusOK, "OK", loginResponse)
}

func (controller *userController) ListAll(w http.ResponseWriter, r *http.Request) {
	const key helper.KeyType = "currentUserID"
	userId := r.Context().Value(key).(string)
	id, _ := primitive.ObjectIDFromHex(userId)

	response, err := controller.UserService.ListAll(id)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	helper.WebResponse(w, http.StatusOK, "OK", response)
}

func (controller *userController) UserDetail(w http.ResponseWriter, r *http.Request) {
	const key helper.KeyType = "currentUserID"
	userId := r.Context().Value(key).(string)
	id, _ := primitive.ObjectIDFromHex(userId)

	response, err := controller.UserService.UserDetail(id)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	helper.WebResponse(w, http.StatusOK, "OK", response)
}

func (controller *userController) Update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	decoder := json.NewDecoder(r.Body)
	user := request.UpdateRequest{}
	err := decoder.Decode(&user)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	const key helper.KeyType = "currentUserID"
	userId := r.Context().Value(key).(string)
	id, _ := primitive.ObjectIDFromHex(userId)

	userResponse, err := controller.UserService.Update(id, username, user)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	helper.WebResponse(w, http.StatusOK, "OK", userResponse)
}

func (controller *userController) Delete(w http.ResponseWriter, r *http.Request) {
	usename := chi.URLParam(r, "username")

	const key helper.KeyType = "currentUserID"
	userId := r.Context().Value(key).(string)
	id, _ := primitive.ObjectIDFromHex(userId)

	userResponse, err := controller.UserService.Delete(id, usename)
	if err != nil {
		helper.WebResponse(w, http.StatusBadRequest, "ERROR", err.Error())
		return
	}

	helper.WebResponse(w, http.StatusOK, "OK", userResponse)
}
