package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// "github.com/imdaaniel/bitcoins-rest-api/api/auth"
	"github.com/imdaaniel/bitcoins-rest-api/api/models"
	"github.com/imdaaniel/bitcoins-rest-api/api/responses"
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/formaterror"
)

func (server *Server) CreateUser(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(res, http.StatusInternalServerError, formattedError)
		return
	}

	res.Header().Set("Location", fmt.Sprintf("%s%s/%d", req.Host, req.RequestURI, userCreated.ID))
	responses.JSON(res, http.StatusCreated, userCreated)
}

func (server *Server) GetUsers(res http.ResponseWriter, req *http.Request) {
	user := models.User{}
	users, err := user.FindUsers(server.DB)

	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusOK, users)
}