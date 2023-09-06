package router

import (
	"Immersive_dash/app/middlewares"
	menteeD "Immersive_dash/features/mentee/data"
	menteeH "Immersive_dash/features/mentee/handler"
	menteeS "Immersive_dash/features/mentee/service"
	userD "Immersive_dash/features/user/data"
	userH "Immersive_dash/features/user/handler"
	userS "Immersive_dash/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := userD.New(db)
	userService := userS.New(userData)
	userHandlerAPI := userH.New(userService)

	menteeData := menteeD.New(db)
	menteeService := menteeS.New(menteeData)
	menteeHandlerAPI := menteeH.New(menteeService)

	// User Endpoint
	e.GET("/users", userHandlerAPI.ReadUser, middlewares.JWTMiddleware())
	e.POST("/login", userHandlerAPI.Login)
	e.POST("/users", userHandlerAPI.RegisterUser, middlewares.JWTMiddleware())

	// Mentee Endpoint
	e.POST("/mentees", menteeHandlerAPI.CreateMentee, middlewares.JWTMiddleware())
}
