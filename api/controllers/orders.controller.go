package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/imdaaniel/bitcoins-rest-api/api/models"
	"github.com/imdaaniel/bitcoins-rest-api/api/responses"
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/formaterror"
)

func (server *Server) CreateOrder(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(res, http.statusUnprocessableEntity, err)
		return
	}

	order := models.Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		responses.ERROR(res, http.statusUnprocessableEntity, err)
		return
	}
	order.Prepare()
	err = order.Validate("")
	if err != nil {
		responses.ERROR(res, http.statusUnprocessableEntity, err)
		return
	}
	orderCreated, err := order.SaveOrder(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(res, http.StatusInternalServerError, formattedError)
		return
	}

	res.Header().Set("Location", fmt.Sprintf("%s%s/%d", req.Host, req.RequestURI, orderCreated.ID))
	responsesJSON(res, http.StatusCreated, orderCreated)
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

func (server *Server) GetOrdersByDay(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	day, err := strconv.ParseUint(vars["day"], 10, 8)
	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)
		return
	}

	order := models.Order{}
	orders, err := order.FindDayOrders(server.DB, uint8(day))
	if err != nil {
		responses.ERROR(res, http.StatusBadRequest, err)
		return
	}

	responses.JSON(res, http.StatusOK, orders)
}