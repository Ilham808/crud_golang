package main

import (
	"OrmGo/config"
	controller "OrmGo/controllers"
	model "OrmGo/models"
	route "OrmGo/routes"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = config.InitConfig()

	db := model.InitModel(*config)
	model.Migrate(db)

	userModel := model.UsersModel{}
	userModel.Init(db)

	bookModel := model.BooksModel{}
	bookModel.Init(db)

	blogModel := model.BlogModel{}
	blogModel.Init(db)

	userControll := controller.UserController{}
	userControll.InitUserController(userModel, *config)

	bookControll := controller.BookController{}
	bookControll.InitBookController(bookModel)

	blogControll := controller.BlogController{}
	blogControll.InitBlogController(blogModel)
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))
	route.RouteUser(e, userControll, *config)
	route.RouteBook(e, bookControll, *config)
	route.RouteBlog(e, blogControll, *config)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", 8000)).Error())
}
