package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/k4utech/m-server/router"
)

func Start() {
	r := rtr.GetRouter()
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}
