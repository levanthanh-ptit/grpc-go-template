package infrastructure

import (
	"grpc-go-templete/internal/user/domain"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddClaims(target jwt.MapClaims, data interface{}) {
	user := data.(*domain.User)
	target["user_id"] = user.ID
}

func ExtractClaims(data jwt.Claims) interface{} {
	user := &domain.User{}
	mapClaims := data.(jwt.MapClaims)
	id, err := primitive.ObjectIDFromHex(mapClaims["user_id"].(string))
	if err != nil {
		return nil
	}
	user.ID = id
	return user
}
