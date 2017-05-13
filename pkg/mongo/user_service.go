package mongo

import (
	"go_web_server/pkg"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
	hash       root.Hash
}

func NewUserService(session *Session, dbName string, collectionName string, hash root.Hash) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, hash}
}

func (us *UserService) Create(u *root.User) error {
	user := newUserModel(u)
	hashedPassword, err := us.hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return us.collection.Insert(&user)
}

func (us *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := us.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}
