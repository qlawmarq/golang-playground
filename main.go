package main

import (
	"net/http"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"playground/libs"
)

	
type commonResponse struct {
    Status int
    Message string
}



func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("/")
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		fmt.Println("/ping")
		return c.JSON(http.StatusOK, commonResponse{Status: 200, Message: "OK"})
	})

	e.GET("/hello", func(c echo.Context) error {
		fmt.Println("/hello")
		libs.Hello()
		libs.HelloForEachNums(2)
		libs.SayGreetingForCurrectTime()
		var num int = 10
		fmt.Printf("%v is bigger than 10: %v\n", num, libs.IsBiggerThan10(num))
		libs.HowToUseVarsInGolang()
		libs.HowToUseDatetime()
		return c.JSON(http.StatusOK, commonResponse{Status: 200, Message: "Hello!"})
	})
		

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}