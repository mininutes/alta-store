package routes

import (
	"alta-store/constraints"
	"alta-store/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/login", controller.LoginUsersController)
	e.POST("/users", controller.CreateUsersController)

	jwt := e.Group("/jwt")
	jwt.Use(middleware.JWT([]byte(constraints.SECRET_JWT)))

	jwt.GET("/carts", controller.GetCartController)
	jwt.POST("/carts", controller.AddToCartController)
	jwt.DELETE("/carts", controller.DeleteCartOnController)

	e.GET("/products", controller.GetProductsController)
	e.GET("/product", controller.GetProductByCategoryController)

	e.GET("/categories", controller.GetCategoriesController)

	return e
}
