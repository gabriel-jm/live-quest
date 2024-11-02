package users

import (
	"errors"
	"log"

	"github.com/labstack/echo/v4"
)

type ReqUser struct {
	Username    string
	DisplayName string
	Email       string
	Picture     string
}

func Routes(e *echo.Echo) *echo.Group {
	group := e.Group("/admin")

	group.POST("/users", func(c echo.Context) error {
		var reqUser ReqUser
		err := c.Bind(&reqUser)

		if err != nil {
			return err
		}

		var repository UserRepository
		user, err := repository.Add(reqUser)

		if err != nil {
			log.Println(err)
			return errors.New("[users] error on create")
		}

		return c.JSON(200, user)
	})

	return group
}
