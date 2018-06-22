package mongo

import (
	"github.com/vodaza36/go-user-mongodb/pck"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserService mongo implementation
type UserService struct {
	collection *mgo.Collection
	hash       root.Hash
}

// NewUserService creates a new instance of the UserService
func NewUserService(session *Session, dbName string, collectionName string, hash root.Hash) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, hash}
}

// CreateUser a new User
func (p *UserService) CreateUser(u *root.User) error {
	user := newUserModel(u)
	hashedPassword, err := p.hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return p.collection.Insert(&user)
}

// GetByUsername returns the root user object, for the given username
func (p *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}
