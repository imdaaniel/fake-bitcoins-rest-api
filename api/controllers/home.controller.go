package controllers

import (
	"net/http"

	"github.com/imdaaniel/bitcoins-rest-api/api/responses"
)

func (server *Server) Home(res http.ResponseWriter, req *http.Request) {
	responses.JSON(res, http.StatusOK, "Be very welcome to the bitcoins API")
}