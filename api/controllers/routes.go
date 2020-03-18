package controllers

import "github.com/imdaaniel/bitcoins-rest-api/api/middlewares"

func (server *Server) InitializeRoutes() {
	// Home
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	// Login
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	// Users
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")

	// Orders
	server.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(server.CreateOrder)).Methods("POST")
	server.Router.HandleFunc("/orders/user/{id}", middlewares.SetMiddlewareJSON(server.GetOrdersByUser)).Methods("GET")
	server.Router.HandleFunc("/orders/date/{date}", middlewares.SetMiddlewareJSON(server.GetOrdersByDate)).Methods("GET")
}
