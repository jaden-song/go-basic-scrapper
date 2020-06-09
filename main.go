package main

import (
	"os"
	"strings"

	"github.com/go-basic-scrapper/scrapper"
	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handlerHome(c echo.Context) error {
	return c.File("home.html")
}

func handlerScrape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handlerHome)
	e.POST("/scrape", handlerScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
