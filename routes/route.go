package route

import (
	"OrmGo/config"
	controller "OrmGo/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc controller.UserController, cfg config.ProgramConfig) {
	e.POST("/login", uc.Login())
	e.GET("/users", uc.Index())
	e.GET("/users/:id", uc.Show())
	e.POST("/users", uc.Create())
	e.PUT("/users/:id", uc.Update(), echojwt.JWT([]byte(cfg.SECRET)))
	e.DELETE("/users/:id", uc.Delete(), echojwt.JWT([]byte(cfg.SECRET)))
}

func RouteBook(e *echo.Echo, uc controller.BookController, cfg config.ProgramConfig) {
	e.GET("/books", uc.Index(), echojwt.JWT([]byte(cfg.SECRET)))
	e.GET("/books/:id", uc.Show(), echojwt.JWT([]byte(cfg.SECRET)))
	e.POST("/books", uc.Create(), echojwt.JWT([]byte(cfg.SECRET)))
	e.PUT("/books/:id", uc.Update())
	e.DELETE("/books/:id", uc.Delete(), echojwt.JWT([]byte(cfg.SECRET)))
}

func RouteBlog(e *echo.Echo, uc controller.BlogController, cfg config.ProgramConfig) {
	e.GET("/blogs", uc.Index())
	e.GET("/blogs/:id", uc.Show())
	e.POST("/blogs", uc.Create())
	e.PUT("/blogs/:id", uc.Update())
	e.DELETE("/blogs/:id", uc.Delete())
}
