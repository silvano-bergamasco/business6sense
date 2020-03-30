package controllers

import "github.com/silvano-bergamasco/business6sense/backend/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	/*s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")*/

	//Companies routes
	s.Router.HandleFunc("/companies", middlewares.SetMiddlewareJSON(s.CreateCompany)).Methods("POST")

	//Stock Markets routes
	s.Router.HandleFunc("/stock_markets", middlewares.SetMiddlewareJSON(s.CreateStockMarket)).Methods("POST")

	//Financial Statements routes
	s.Router.HandleFunc("/financial_statements", middlewares.SetMiddlewareJSON(s.CreateFinancialStatement)).Methods("POST")

	//Stocks routes
	s.Router.HandleFunc("/stocks", middlewares.SetMiddlewareJSON(s.CreateStock)).Methods("POST")

	//Portfolios routes
	s.Router.HandleFunc("/portfolios", middlewares.SetMiddlewareJSON(s.CreatePortfolio)).Methods("POST")

}
