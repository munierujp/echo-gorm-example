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

type UserRequest struct {
	Name       string `json:"name"`
	LanguageID uint   `json:"language_id"`
}

func (h *Handler) AddUser(c echo.Context) error {
	req := new(UserRequest)
	if err := c.Bind(req); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Invalid request",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	user := &model.User{
		Name:       req.Name,
		LanguageID: req.LanguageID,
	}

	userRepo := database.NewUserRepository(h.DB)
	if err := userRepo.Add(user); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
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
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
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
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Not found user",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Invalid ID",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	req := new(UserRequest)
	if err := c.Bind(req); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Invalid request",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	user := &model.User{}
	user.ID = id
	user.Name = req.Name
	user.LanguageID = req.LanguageID

	userRepo := database.NewUserRepository(h.DB)
	if err := userRepo.Update(user); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Failed to update",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Invalid ID",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	userRepo := database.NewUserRepository(h.DB)
	if err := userRepo.Delete(id); err != nil {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/echo-gorm-example/blob/master/handler/user.go",
			Title: "Failed to delete",
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	return c.NoContent(http.StatusNoContent)
}
