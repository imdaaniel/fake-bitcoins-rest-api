package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/imdaaniel/bitcoins-rest-api/api/auth"
	"github.com/imdaaniel/bitcoins-rest-api/api/models"
	"github.com/imdaaniel/bitcoins-rest-api/api/responses"
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/bitcoin"
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/date"
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/formaterror"
)

func (server *Server) CreateOrder(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}

	order := models.Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(req)
	if err != nil {
		responses.ERROR(res, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	tokenIDint, err := strconv.ParseUint(tokenID, 10, 64)
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	if tokenIDint != order.AuthorID {
		responses.ERROR(res, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	order.Prepare()

	order.Value = bitcoin.AmountToValue(order.Amount)

	err = order.Validate()
	if err != nil {
		responses.ERROR(res, http.StatusUnprocessableEntity, err)
		return
	}
	orderCreated, err := order.SaveOrder(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(res, http.StatusInternalServerError, formattedError)
		return
	}

	res.Header().Set("Location", fmt.Sprintf("%s%s/%d", req.Host, req.RequestURI, orderCreated.ID))
	responses.JSON(res, http.StatusCreated, orderCreated)
}

func (server *Server) GetOrdersByUser(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	userID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)
		return
	}

	order := models.Order{}
	orders, err := order.FindUserOrders(server.DB, uint64(userID))
	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusOK, orders)
}

func (server *Server) GetOrdersByDate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	getDate := vars["date"]

	if date.ValiDate(getDate) == false {
		responses.ERROR(res, http.StatusExpectationFailed, errors.New("Invalid Date"))
		return
	}

	order := models.Order{}
	orders, err := order.FindDayOrders(server.DB, getDate)
	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusOK, orders)
}
