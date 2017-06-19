package server

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/k4utech/m-server/router"
	"github.com/k4utech/m-server/services/mongo-db-service"
)

//Start : Starts the server and make sure
// 1. DB Connection set up if not create panic
// 2. Set up all routes
// 3. If all goes well start the server
func Start() {
	mdbs.Init()
	r := rtr.GetRouter()
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, r))

}

func initdb() {

}
