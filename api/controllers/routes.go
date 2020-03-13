package controllers

import "github.com/imdaaniel/bitcoins-rest-api/api/middlewares"

func (server *Server) initializeRoutes() {
	// Home
	server.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(server.Home)).Methods("GET")

	// Login
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(server.Login)).Methods("POST")

	// Users
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.CreateUser)).Methods("POST")

	// Orders
	server.Router.HandleFunc("/orders", middlewares.SetMiddlewareJSON(server.CreteOrder)).Methods("POST")
	server.Router.HandleFunc("/orders/user/{id}", middlewares.SetMiddlewareJSON(server.GetOrdersByUser)).Methods("GET")
	server.Router.HandleFunc("/orders/{day}", middlewares.SetMiddlewareJSON(server.GetOrdersByDay)).Methods("GET")
}