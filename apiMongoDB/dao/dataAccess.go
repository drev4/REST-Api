package dao

import (
	"log"

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

func (m *SportsDAO) FindAll() ([]Sport, error) {
	var deportes []Sport
	err := db.C(COLLECTION).Find(bson.M{}).All(&deportes)
	return deportes, err
}

func (m *SportsDAO) FindById(id string) (Sport, error) {
	var deporte Sport
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&deporte)
	return deporte, err
}

func (m *SportsDAO) Insert(deporte Sport) error {
	err := db.C(COLLECTION).Insert(&deporte)
	return err
}

func (m *SportsDAO) Delete(deporte Sport) error {
	err := db.C(COLLECTION).Remove(&deporte)
	return err
}

func (m *SportsDAO) Update(deporte Sport) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &deporte)
	return err
}
