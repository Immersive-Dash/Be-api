package main

import (
	"Immersive_dash/app/config"
	"Immersive_dash/app/database"
	"Immersive_dash/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbMysql := database.InitDBMysql(cfg)
	database.InitialMigration(dbMysql)

	// create a new echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	router.InitRouter(dbMysql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":443"))
}
