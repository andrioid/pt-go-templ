package main

import (
	"app/web/pages"

	"github.com/labstack/echo/v4"
)

var FormList []string = []string{"Initial item"}

func main() {

	e := echo.New()
	// TODO: Serve /public for CSS
	// TODO: Embed assets into binary
	// TODO: Experiment with SQLite, maybe Gorm
	// TODO: Live reload with Air maybe

	e.GET("/", func(c echo.Context) error {
		return pages.RootIndexPage(FormList).Render(c.Request().Context(), c.Response().Writer)
		//return Hello("hello").Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("/", func(c echo.Context) error {
		ctx := c.Request().Context()

		FormList = append(FormList, c.FormValue("add-item"))
		c.Response().Header().Set("Content-Type", "text/html")
		return pages.Counts(FormList).Render(ctx, c.Response().Writer)
	})
	e.Static("/", "./web/public")
	e.Logger.Fatal(e.Start(":1323"))
}
