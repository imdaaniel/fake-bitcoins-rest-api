package controllers

import (
	"net/http"

	"github.com/imdaaniel/bitcoins-rest-api/api/responses"
)

func (server *Server) Home(res http.ResponseWriter, req *http.Request) {
	responses.JSON(res, http.StatusOK, map[string]string{
		"message": "Be very welcome to my bitcoins API",
	})
}
