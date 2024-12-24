package middleware

import (
	"context"
	"session-22/jwt"
	"session-22/model"
	"session-22/util"

	"net/http"
)

const (
	AuthHeader = "Authorization"
)

func CheckAuth(next http.HandlerFunc) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get(AuthHeader)

		if token == "" {
			util.PrepareResponse(w, "", &model.UnAuthorized)
			return
		}

		username, err := jwt.Verify(token)

		if err != nil {

			util.PrepareResponse(w, "", &model.UnAuthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "USER_NAME", username)
		r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})

}
