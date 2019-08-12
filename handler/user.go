package handler

import (
	"echo-gorm-example/domain/model"
	"echo-gorm-example/domain/model/response"
	"echo-gorm-example/interface/database"
	"echo-gorm-example/modules"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) AddUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/users.go",
			Title: "Invalid request",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	userRepo := database.NewUserRepository(h.DB)
	if err := userRepo.Add(user); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/users.go",
			Title: "Failed to create",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetUser(c echo.Context) error {
	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/users.go",
			Title: "Invalid ID",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	userRepo := database.NewUserRepository(h.DB)
	user, err := userRepo.FindByID(id)
	if err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/users.go",
			Title: "Not found user",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	return c.JSON(http.StatusOK, user)
}

// TODO: UpdateUser

func (h *Handler) DeleteUser(c echo.Context) error {
	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/users.go",
			Title: "Invalid ID",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	userRepo := database.NewUserRepository(h.DB)
	if err := userRepo.Delete(id); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/users.go",
			Title: "Failed to delete",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	return c.NoContent(http.StatusNoContent)
}
