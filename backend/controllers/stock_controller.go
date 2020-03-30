package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"github.com/silvano-bergamasco/business6sense/backend/auth"
	"github.com/silvano-bergamasco/business6sense/backend/models"
	"github.com/silvano-bergamasco/business6sense/backend/responses"
	"github.com/silvano-bergamasco/business6sense/backend/utils/formaterror"
)

func (server *Server) CreateStock(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	s := models.Stock{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = s.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = s.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	/*uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}*/

	/*if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}*/

	err = s.InsertStock()
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	var stocks []models.Stock

	stocks = append(stocks, s)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, s.ID))
	responses.JSON(w, http.StatusCreated, stocks)
}
