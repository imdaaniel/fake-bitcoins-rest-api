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
	"github.com/imdaaniel/bitcoins-rest-api/api/utils/formaterror"
)