package application

import (
	"context"
	"grpc-go-templete/pkg/http_utils"
	"grpc-go-templete/pkg/pb/user_pb"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
)

var (
	_unauthenticatedError = http_utils.Error{
		Code:    uint32(codes.Unauthenticated),
		Message: "UNAUTHENTICATED",
	}
	_authenticationFailedError = http_utils.Error{
		Code:    uint32(codes.Unavailable),
		Message: "AUTHENTICATION_FAILED",
	}
)

func (productsServer *productsGrpcServer) AuthGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http_utils.JSONError(w, _unauthenticatedError, http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http_utils.JSONError(w, _unauthenticatedError, http.StatusUnauthorized)
			return
		}
		data, err := productsServer.UsersClient.VerifyAuthToken(context.Background(), &user_pb.VerifyAuthTokenRequest{Token: token})
		if err != nil || data == nil {
			http_utils.JSONError(w, _authenticationFailedError, http.StatusBadGateway)
			return
		}
		if !data.Authenticated {
			http_utils.JSONError(w, _unauthenticatedError, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
