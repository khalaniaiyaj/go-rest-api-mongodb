package dao

import (
	. "github.com/user/golang-new/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Dao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	MOVIE_COLLECTION = "movies"
	USER_COLLECTION  = "user"
)

func (m *Dao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *Dao) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(MOVIE_COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *Dao) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(MOVIE_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *Dao) Insert(movie Movie) error {
	err := db.C(MOVIE_COLLECTION).Insert(&movie)
	return err
}

func (m *Dao) Delete(movie Movie) error {
	err := db.C(MOVIE_COLLECTION).Remove(&movie)
	return err
}

func (m *Dao) Update(movie Movie) error {
	err := db.C(MOVIE_COLLECTION).UpdateId(movie.ID, &movie)
	return err
}

func (m *Dao) CreateUser(user User) error {
	err := db.C(USER_COLLECTION).Insert(&user)
	return err
}

func (m *Dao) FindUserByName(name string) (User, error) {
	var user User
	err := db.C(USER_COLLECTION).Find(bson.M{"name": name}).One(&user)
	return user, err
}
