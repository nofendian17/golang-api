package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nofendian17/golang-api/entity"
	"github.com/nofendian17/golang-api/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entity.User {
	emptyData := make([]entity.User, 0)
	if len(c.service.FindAll()) > 0 {
		return c.service.FindAll()
	} else {
		return emptyData
	}
}

func (c controller) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}

func (c controller) Update(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	user.ID = oid
	c.service.Update(user)
	return nil
}

func (c controller) Delete(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	id := ctx.Param("id")
	oid, err := primitive.ObjectIDFromHex(id)
	user.ID = oid
	c.service.Delete(user)
	return nil
}
