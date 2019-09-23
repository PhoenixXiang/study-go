package main

import (
	"fmt"
	"io"
	"os"
	"path"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	p := "./temp/" + file.Filename
	err = os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}
	dst, err := os.Create(p)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := e.Group("/test")
	t.Static("/", "public")
	t.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}
