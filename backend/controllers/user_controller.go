package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/silvano-bergamasco/business6sense/backend/models"
	"github.com/silvano-bergamasco/business6sense/backend/responses"
	"github.com/silvano-bergamasco/business6sense/backend/utils/dbutils"
	"github.com/silvano-bergamasco/business6sense/backend/utils/formaterror"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	u := models.User{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = u.Prepare(dbutils.Insert)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = u.Validate(dbutils.Insert)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = u.InsertUser()
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	var users []models.User

	users = append(users, u)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, u.ID))
	responses.JSON(w, http.StatusCreated, users)
}
