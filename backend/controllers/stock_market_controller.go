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

func (server *Server) CreateStockMarket(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	m := models.StockMarket{}
	err = json.Unmarshal(body, &m)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = m.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = m.Validate()
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

	m.Active = true

	err = m.InsertMarket()
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	var markets []models.StockMarket

	markets = append(markets, m)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, m.ID))
	responses.JSON(w, http.StatusCreated, markets)
}
