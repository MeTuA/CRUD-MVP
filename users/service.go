package users

import (
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
)

type service struct {
	store Store
	log   hclog.Logger
}

func NewService(store Store, log hclog.Logger) *echo.Echo {

	s := &service{
		store: store,
		log:   log,
	}
	app := echo.New()
	app.GET("/api/getUsers", s.GetAllUsers)
	app.POST("/api/createUser", s.CreateUser)
	app.POST("/api/deleteUser/:id", s.DeleteUser)
	app.POST("/api/updateUser/:id", s.UpdateUser)

	return app
}

func (s *service) GetAllUsers(c echo.Context) error {
	users, err := s.store.GetAllUsers()
	if err != nil {
		s.log.Error("failed on getting all users", "error", err.Error())
		return c.NoContent(500)
	}

	return c.JSON(200, users)
}

func (s *service) CreateUser(c echo.Context) error {
	user := &User{}

	err := c.Bind(&user)
	if err != nil {
		s.log.Error("failed to bind in CreateUser", "error", err.Error())
		return c.JSON(400, err.Error())
	}

	err = s.store.CreateUser(user)
	if err != nil {
		s.log.Error("failed on creating user", "error", err.Error())
		return c.JSON(500, err.Error())
	}

	return c.NoContent(200)
}

func (s *service) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := s.store.DeleteUser(id)
	if err != nil {
		s.log.Error("failed to delete user", "error", err.Error())
		return c.JSON(500, err.Error())
	}

	return c.NoContent(200)
}

func (s *service) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := &User{}

	err := c.Bind(&user)
	if err != nil {
		s.log.Error("failed to bind in UpdateUser", "error", err.Error())
		return c.JSON(400, err.Error())
	}

	err = s.store.UpdateUser(id, user)
	if err != nil {
		s.log.Error("failed to delete user", "error", err.Error())
		return c.JSON(500, err.Error())
	}

	return c.NoContent(200)
}
