package application

import (
	"context"
	"errors"
	"grpc-go-templete/pkg/pb/user_pb"
	"net/http"
	"strings"

	"github.com/levanthanh-ptit/go-ez/ez_grpc"
	"github.com/levanthanh-ptit/go-ez/ez_http"
)

var (
	_unauthenticatedError      = ez_grpc.MakeUnauthenticated(errors.New("UNAUTHENTICATED"))
	_authenticationFailedError = ez_grpc.MakeUnavailable(errors.New("AUTHENTICATION_FAILED"))
)

func (s *GrpcGatewayServer) AuthGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ez_http.JSONError(w, ez_grpc.ToErrorJSONStatus(_unauthenticatedError), http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			ez_http.JSONError(w, ez_grpc.ToErrorJSONStatus(_unauthenticatedError), http.StatusUnauthorized)
			return
		}
		data, err := s.usersClient.VerifyAuthToken(context.Background(), &user_pb.VerifyAuthTokenRequest{Token: token})
		if err != nil || data == nil {
			ez_http.JSONError(w, ez_grpc.ToErrorJSONStatus(_authenticationFailedError), http.StatusBadGateway)
			return
		}
		if !data.Authenticated {
			ez_http.JSONError(w, ez_grpc.ToErrorJSONStatus(_unauthenticatedError), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
