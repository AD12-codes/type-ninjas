package users

import (
	"github.com/AD12-codes/go-template/db"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, q *db.Queries) {
	s := NewService(q)
	h := NewHandler(s)

	group := e.Group("api/v1/users")
	group.POST("", h.CreateUserHandler)
	group.GET("", h.GetAllUsersHandler)
}
