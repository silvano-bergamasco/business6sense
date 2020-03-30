package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/silvano-bergamasco/business6sense/backend/models"
	"github.com/silvano-bergamasco/business6sense/backend/responses"
	"github.com/silvano-bergamasco/business6sense/backend/utils/formaterror"
)

func (server *Server) CreateFinancialStatement(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	fs := models.FinancialStatement{}
	err = json.Unmarshal(body, &fs)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = fs.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = fs.Validate()
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

	err = fs.InsertFinancialStatement()
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	var finantialStatements []models.FinancialStatement

	finantialStatements = append(finantialStatements, fs)

	w.Header().Set("Location", fmt.Sprintf("%s%s/%s", r.Host, r.URL.Path, fs.ID))
	responses.JSON(w, http.StatusCreated, finantialStatements)
}
