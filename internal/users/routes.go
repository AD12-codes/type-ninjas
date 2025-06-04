package users

import (
	"github.com/AD12-codes/type-ninjas/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group, queries *db.Queries) {
	s := NewService(queries)
	h := NewHandler(s)

	userGroup := e.Group("/users")
	userGroup.POST("/register", h.RegisterUserHandler)
	userGroup.GET("/me", h.MeHandler)
	userGroup.GET("", h.GetAllUsersHandler)
	userGroup.GET("/:id", h.GetUserHandler)

}
