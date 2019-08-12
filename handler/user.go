package handler

import (
	"echo-gorm-example/domain/model/response"
	"echo-gorm-example/interface/database"
	"echo-gorm-example/modules"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: AddUser

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

// TODO: DeleteUser
