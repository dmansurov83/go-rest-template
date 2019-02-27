package main

import (
	"./env"
	"./services"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type Server struct {
	config        *env.Config
	personService services.PersonService
}

func bindPersons(e *echo.Echo, personService services.PersonService) {
	e.GET("/persons", func(c echo.Context) error {
		persons, err := personService.FindAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, persons)
	})
	e.GET("/persons/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		person, err := personService.Get(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, person)
	})
}

func (s *Server) Run() {
	e := echo.New()
	bindPersons(e, s.personService)
	e.Logger.Fatal(e.Start(":" + s.config.Port))
}

func NewServer(config *env.Config,
	personService services.PersonService) *Server {
	return &Server{
		config,
		personService,
	}
}
