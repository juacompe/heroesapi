package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/heroes", listHeroes)

	e.Logger.Fatal(e.Start(":1323"))
}

type Hero struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func listHeroes(c echo.Context) error {
	hs := []Hero{
		Hero{11, "Mr. Nice"},
		{12, "Narco"},
		{13, "Bombasto"},
		{14, "Celeritas"},
		{15, "Magneta"},
		{16, "RubberMan"},
		{17, "Dynama"},
		{18, "Dr IQ"},
		{19, "Magma"},
		{20, "Tornado"},
	}
	return c.JSON(http.StatusOK, hs)
}
