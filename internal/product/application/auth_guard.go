package application

import (
	"context"
	"grpc-go-templete/pkg/pb/user_pb"
	"net/http"
	"strings"
)

func (productsServer *productsGrpcServer) AuthGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" && !strings.HasPrefix(authHeader, "Bearer ") {
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			return
		}
		data, err := productsServer.UsersClient.VerifyAuthToken(context.Background(), &user_pb.VerifyAuthTokenRequest{Token: token})
		if err != nil {
			http.Error(w, "AUTHENTICATION_FAILED", http.StatusBadGateway)
			return
		}
		if !data.Authenticated {
			http.Error(w, "UNAUTHENTICATED", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
