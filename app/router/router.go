package router

import (
	"Immersive_dash/features/user/data"
	"Immersive_dash/features/user/handler"
	"Immersive_dash/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := data.New(db)
	userService := service.New(userData)
	userHandlerAPI := handler.New(userService)

	e.GET("/users", userHandlerAPI.ReadUser)
}
