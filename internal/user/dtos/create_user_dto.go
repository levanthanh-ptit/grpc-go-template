package dtos

type CreateUserDto struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
