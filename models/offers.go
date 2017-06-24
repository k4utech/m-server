package mdl

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/k4utech/m-server/services/mongo-db-service"
)

//Offer :Model to represent user
type Offer struct {
	Imagesrc string `json:"imagesrc"`
	Header   string `json:"header"`
	Content  string `json:"content"`
}

//GetAll : Get all the offers
//Not being used till now
//USe as examples
func (*Offer) GetAll() {
	fmt.Println("getting offers..")
	var offers []Offer
	mdbs.DBConn.C("notifications").Find(bson.M{}).All(&offers)
	fmt.Println(offers)
}
