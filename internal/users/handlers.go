package users

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterUserHandler(c echo.Context) error {
	var req RegisterUserRequest

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid request format"})
	}

	if err := c.Validate(&req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		messages := make([]string, 0)
		for _, v := range validationErrors {
			messageStr := fmt.Sprintf("Field '%s' failed on '%s' rule", v.Field(), v.Tag())
			messages = append(messages, messageStr)
		}
		return c.JSON(http.StatusBadRequest, echo.Map{"message": strings.Join(messages, ", ")})
	}
	registeredUserError := h.service.RegisterUser(c.Request().Context(), req)

	if registeredUserError != nil {
		fmt.Println(registeredUserError)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": registeredUserError.Error()})
	}

	return c.NoContent(http.StatusCreated)
}

// role: admin | scope: read:usersAllData
// role: user not allowed
func (h *Handler) GetAllUsersHandler(c echo.Context) error {
	users, err := h.service.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

// role: admin | scope: read:usersAllData
// role: user | scope: read:usersSelfData || read:usersLimitedData (username, joined date)
func (h *Handler) GetUserHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid user ID"})
	}
	user, err := h.service.GetUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

// role: any | scope: read:usersSelfData
// TODO: implement meHandler with rbac
func (h *Handler) MeHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, c.Get("user"))
}
