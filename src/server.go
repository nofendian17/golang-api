package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nofendian17/logAudit/controller"
	"github.com/nofendian17/logAudit/middlewares"
	"github.com/nofendian17/logAudit/repository"
	"github.com/nofendian17/logAudit/service"
	"net/http"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository()
	userService service.UserService = service.New(userRepository)
	userController controller.UserController = controller.New(userService)
)


func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Log())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app_name": "log audit golang",
			"version": "1.0.0",
			"message": "welcome to application",
		})
	})

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.FindAll())
	})

	server.POST("/users", func(ctx *gin.Context) {
		err := userController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		} else {
			ctx.JSON(200, userController.FindAll())
		}

	})

	server.PUT("/users/:id", func(ctx *gin.Context) {
		err := userController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		} else {
			ctx.JSON(200, userController.FindAll())
		}

	})

	server.DELETE("/users/:id", func(ctx *gin.Context) {
		err := userController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		} else {
			ctx.JSON(200, userController.FindAll())
		}

	})
	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}