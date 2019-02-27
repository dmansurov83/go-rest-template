package main

import (
	"./env"
	"./services"
	"github.com/labstack/echo"
	"net/http"
)

type Server struct {
	config        *env.Config
	personService services.PersonService
}

func (s *Server) Run() {
	e := echo.New()
	e.GET("/persons", func(c echo.Context) error {
		return c.JSON(http.StatusOK, s.personService.FindAll())
	})
	e.Logger.Fatal(e.Start(":" + s.config.Port))
}

func NewServer(config *env.Config,
	personService services.PersonService) *Server {
	return &Server{
		config,
		personService,
	}
}
