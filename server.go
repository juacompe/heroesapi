package main

import (
	"net/http"

	"github.com/juacompe/heroesapi/hero"
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

func listHeroes(c echo.Context) error {
	hs := []hero.Hero{
		{11, "Mr. Nice"},
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
