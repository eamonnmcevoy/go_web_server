package mongo

import (
	"go_web_server/pkg"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userModel struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Password string
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUserModel(u *root.User) *userModel {
	return &userModel{
		Username: u.Username,
		Password: u.Password}
}

func (u *userModel) toRootUser() *root.User {
	return &root.User{
		Id:       u.ID.Hex(),
		Username: u.Username,
		Password: u.Password}
}
