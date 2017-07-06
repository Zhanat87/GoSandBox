package SessionManager

import (
	"encoding/json"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gorilla/sessions"
	Password "github.com/nkozyra/api/password"
	"time"
)

var Session UserSession

type UserSession struct {
	ID              string            `json:"id"`
	GorillaSesssion *sessions.Session `json:"session"`
	SessionStore    *memcache.Client  `json:"store"`
	UID             int               `json:"uid"`
	Expire          time.Time         `json:"expire"`
}

func (us *UserSession) Create() {
	us.SessionStore = memcache.New("127.0.0.1:11211")
	us.ID = Password.GenerateSessionID(32)
}

func (us *UserSession) GetSession(key string) (UserSession, error) {
	session, err := us.SessionStore.Get(us.ID)
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
	us.SessionStore.Set(&memcache.Item{Key: us.ID, Value: []byte(jsonValue)})
	_, err := us.SessionStore.Get(us.ID)
	if err != nil {
		return false
	} else {
		return true
	}
}
