package repository

import (
	"fmt"
	"github.com/Kamva/mgm/v3"
	"github.com/nofendian17/golang-api/db"
	"github.com/nofendian17/golang-api/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	Save(user entity.User)
	Update(user entity.User)
	Delete(user entity.User)
	FindAll() []entity.User
}

type database struct {
	connection *mgm.Collection
}

func NewUserRepository() UserRepository {
	db.ConnectMongoDB()
	coll := &entity.User{}
	dbs := mgm.Coll(coll)
	return &database{connection: dbs}
}

func (db *database) Save(user entity.User) {
	_ = db.connection.Create(&user)
}

func (db *database) Update(user entity.User) {
	_ = db.connection.Update(&user)
}

func (db *database) Delete(user entity.User) {
	_ = db.connection.Delete(&user)
}

func (db *database) FindAll() []entity.User {
	var users []entity.User
	err := db.connection.SimpleFind(&users, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	return users
}
