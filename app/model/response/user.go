package response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	Id       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
}

type LoginResponse struct {
	Id       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
	Token    string             `json:"token"`
}

type ShowAll struct {
	Id       primitive.ObjectID `json:"_id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
}
