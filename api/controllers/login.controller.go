package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/imdaaniel/bitcoins-rest-api/api/auth"
	"github.com/imdaaniel/bitcoins-rest-api/api/models"
	"github.com/imdaaniel/bitcoins-rest-api/api/responses"
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.JSON(res, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.JSON(res, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.JSON(res, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.Formaterror(err.Error())
		responses.JSON(res, http.StatusUnprocessableEntity, formattedError)
		return
	}

	responses.JSON(res, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}