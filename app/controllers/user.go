package controllers

import (
	"github.com/dbenson24/revel-testing/mongodb"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	*revel.Controller
	mongodb.MongoController
}

type person struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Phone string `json:"phone" bson:"phone"`
	Age   int    `json:"age" bson:"age"`
}

func (u *User) Index() revel.Result {

	var c *mgo.Collection
	c = u.Mongo.DB("revel-test").C("users")
	var results []person
	_ = c.Find(bson.M{}).All(&results)
	return u.Render(results)
}
