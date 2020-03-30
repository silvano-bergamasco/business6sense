package controllers

import (
	"net/http"

	"github.com/silvano-bergamasco/business6sense/backend/responses"
)

// Home: Healthcheck endpoint
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Business6sense API")
}
