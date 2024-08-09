package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Person struct {
	Name     string
	Lastname string
}

func newPerson(name, lastname string) Person {
	return Person{
		Name:     name,
		Lastname: lastname,
	}
}

type People = []Person

// data for site
type Data struct {
	People People
}

func newData() Data {
	return Data{
		People: []Person{
			newPerson("Filbert", "Green"),
			newPerson("Gregorny", "Hilbert"),
		},
	}
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	data := newData()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", data)
	})

	e.POST("/people", func(c echo.Context) error {
		name := c.FormValue("name")
		lastname := c.FormValue("lastname")

		data.People = append(data.People, newPerson(name, lastname))
		return c.Render(200, "personlist", data)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
