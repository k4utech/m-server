package dao

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/k4utech/m-server/services/mongo-db-service"

	"github.com/k4utech/m-server/models"
)

func GetAllOffers() []mdl.Offer {
	fmt.Printf("Featching all offers from db...")
	var offers []mdl.Offer
	mdbs.DBConn.C("notifications").Find(bson.M{}).All(&offers)
	fmt.Println(offers)
	return offers
}
