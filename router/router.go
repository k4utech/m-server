package rtr

import (
	"github.com/gorilla/mux"
	"github.com/k4utech/m-server/handlers"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/v1/get-token", hdlrs.GetTokenHandler).Methods("GET")
	r.Handle("/v1/status", hdlrs.StatusHandler).Methods("GET")
	r.Handle("/v1/products", hdlrs.ProductsHandler).Methods("GET")
	return r
}
