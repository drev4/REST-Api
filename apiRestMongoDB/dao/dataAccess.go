package dao

import (
	"log"
	"apiRestMongoDB/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SportsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "deportes"
)

func (m *SportsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *SportsDAO) FindAll() ([]model.Sport, error) {
	var deportes []model.Sport
	err := db.C(COLLECTION).Find(bson.M{}).All(&deportes)
	return deportes, err
}

func (m *SportsDAO) FindById(id string) (model.Sport, error) {
	var deporte model.Sport
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&deporte)
	return deporte, err
}

func (m *SportsDAO) Insert(deporte model.Sport) error {
	err := db.C(COLLECTION).Insert(&deporte)
	return err
}

func (m *SportsDAO) Delete(deporte model.Sport) error {
	err := db.C(COLLECTION).Remove(&deporte)
	return err
}

func (m *SportsDAO) Update(deporte model.Sport) error {
	err := db.C(COLLECTION).UpdateId(deporte.ID, &deporte)
	return err
}


