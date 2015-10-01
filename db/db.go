package db

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/revel/revel"
)

const (
	mongoDBHosts = "ds057963.mongolab.com:57963"
	authDatabase = "revel-test"
	authUserName = "test"
	authPassword = "test"
)

// Variable representing the master db session
var Db *mgo.Session

// MongoController Adds a Mongo Session pointer to the controller
type MongoController struct {
	*revel.Controller
	Session *mgo.Session
}

// Open provides a MongoController with a copy of the master session
func (c *MongoController) Open() revel.Result {
	c.Session = Db.Copy()
	return nil
}

// Close closes the copy after the request is made
func (c *MongoController) Close() revel.Result {
	c.Session.Close()
	return nil
}

// Error closes the session in case of a panic
func (c *MongoController) Error() revel.Result {
	c.Session.Close()
	return nil
}

// InitDB opens an initial connection to the Database.
func InitDB() {
	log.Println("InitDB")
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{mongoDBHosts},
		Timeout:  30 * time.Second,
		Database: authDatabase,
		Username: authUserName,
		Password: authPassword,
	}
	Db, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	Db.SetMode(mgo.Monotonic, true)
	//bson.NewObjectId()
	revel.InterceptMethod((*MongoController).Open, revel.BEFORE)
	revel.InterceptMethod((*MongoController).Close, revel.AFTER)
	revel.InterceptMethod((*MongoController).Error, revel.PANIC)
	log.Println("Executed")
}
