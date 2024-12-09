package user

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/romulosm/go-modular-backend-template/internal/user/handler"
	"github.com/romulosm/go-modular-backend-template/internal/user/repository"
	"github.com/romulosm/go-modular-backend-template/internal/user/service"
)

func InitModule(r *gin.Engine, db *sql.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userHandler.CreateUser)
		userRoutes.GET("/:id", userHandler.GetUser)
	}
}
