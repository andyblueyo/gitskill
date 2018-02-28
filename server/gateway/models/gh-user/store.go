package gh_user

import (
	"gopkg.in/mgo.v2/bson"
)

type Store interface {
	GetByID(id bson.ObjectId) (*User, error)

	GetByGithubName(name string) (*User, error)

	Insert(user *User) error
	//Update(memeID bson.ObjectId, updates *MemeUpdates, userID bson.ObjectId) (*Meme, error)
}
