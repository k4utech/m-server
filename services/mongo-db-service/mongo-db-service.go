package mdbs

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// MongoDbSession : new type to represent session used by mongo
type MongoDbSession *mgo.Session

// DBSession : this will be exposed to other packages to access mongo db session
var DBSession *mgo.Session

//DBConn : db to connect to collections
var DBConn *mgo.Database

//Init : Perform the following tasks
// Set up a connection with mongo server
// return session and err as per connection result
func Init() (MongoDbSession, error) {
	session, err := mgo.Dial("path")
	if err != nil {
		fmt.Println("Error while connecting to db server")
		fmt.Println(err.Error())
		panic("Error while connecting to db server")
	} else {
		fmt.Println("Connected to db")
		DBSession = session
		DBConn = DBSession.DB("gmart_prod")
	}
	return DBSession, err

}

//Destroy : Destroy db connection
func Destroy() {
	fmt.Println("Destroying Db Connection...")
	DBSession.Close()
}
