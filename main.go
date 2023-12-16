package main

import (
	"app/repos/todo"
	"app/web/pages"
	"database/sql"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func main() {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	DB = db

	e := echo.New()
	// TODO: Serve /public for CSS
	// TODO: Embed assets into binary
	// TODO: Experiment with SQLite, maybe Gorm
	// TODO: Live reload with Air maybe

	e.GET("/", func(c echo.Context) error {
		return pages.RootIndexPage().Render(c.Request().Context(), c.Response().Writer)
		//return Hello("hello").Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("/", func(c echo.Context) error {
		ctx := c.Request().Context()
		list := todo.AddItem(c.FormValue("add-item"))

		c.Response().Header().Set("Content-Type", "text/html")
		return pages.Counts(list).Render(ctx, c.Response().Writer)
	})
	e.Static("/", "./web/public")
	e.Logger.Fatal(e.Start(":1323"))
}
