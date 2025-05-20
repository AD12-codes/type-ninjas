package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateUserHandler(c echo.Context) error {
	var req CreateUserRequest

	if err := c.Bind(&req); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, echo.Map{"err": err})
	}

	createUserError := h.service.CreateUser(c.Request().Context(), req)

	if createUserError != nil {
		log.Print(createUserError)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": createUserError.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAllUsersHandler(c echo.Context) error {
	users, err := h.service.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}
