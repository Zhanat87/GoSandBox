package SessionManager

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/sessions"
	Password "github.com/nkozyra/api/password"
	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var Session UserSession

type UserSession struct {
	ID              string            `bson:"_id"`
	GorillaSesssion *sessions.Session `bson:"session"`
	SessionStore    *mgo.Collection   `bson:"store"`
	UID             int               `bson:"uid"`
	Value           []byte            `bson:"Valid"`
	Expire          time.Time         `bson:"expire"`
}

func (us *UserSession) Create() {
	s, err := mgo.Dial("127.0.0.1:27017/sessions")
	defer s.Close()
	if err != nil {
		log.Println("Can't connect to MongoDB")
	} else {
		us.SessionStore = s.DB("sessions").C("sessions")
	}
	us.ID = Password.GenerateSessionID(32)
}

func (us *UserSession) GetSession(key string) (UserSession, error) {
	var session UserSession
	err := us.SessionStore.Find(us.ID).One(session)
	if err != nil {
		return UserSession{}, errors.New("No such session")
	} else {
		var tempSession = UserSession{}
		err := json.Unmarshal(session.Value, tempSession)
		if err != nil {

		}
		return tempSession, nil
	}
}

func (us *UserSession) SetSession() bool {
	jsonValue, _ := json.Marshal(us)
	err := us.SessionStore.Insert(UserSession{ID: us.ID, Value: []byte(jsonValue)})
	if err != nil {
		return false
	} else {
		return true
	}
}
