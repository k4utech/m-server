package hdlrs

import (
	"encoding/json"
	"net/http"

	"github.com/k4utech/m-server/daos"
	"github.com/k4utech/m-server/models"
)

var offers = []mdl.Offer{}

//OfferHandler : Handles offer endpoint request
var OfferHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	offers = dao.GetAllOffers()
	payload, _ := json.Marshal(offers)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(payload))
})
