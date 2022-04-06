package auth

import "github.com/gorilla/mux"

func RegisterRoutes(route *mux.Router, service *Service) {
	route.HandleFunc("/users/register", registerUser(service))
	route.HandleFunc("/users/auth", authUser(service))
}
