package middlewares

import (
	"errors"
	"net/http"

	"github.com/imdaaniel/rest-api/api/auth"
	"github.com/imdaaniel/rest-api/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next(res, req)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		err := auth.TokenValid(req)
		if err != nil {
			responses.ERROR(res, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(res, req)
	}
}