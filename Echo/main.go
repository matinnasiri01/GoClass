package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

//----------------------------
func WelcomeHandler(c echo.Context) error {
	return c.String(http.StatusOK,"Welcome!")
}

//----------------------------
type User struct{
	Username string `json:"username"`
	Age int		 `json:"age"`
}
func SignupHandler(c echo.Context) error {
	user := new(User)
	if err:= c.Bind(user);err != nil {
		return err
	}
	return c.JSON(http.StatusCreated,user)
}
func SayHiHandler(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.String(http.StatusBadRequest, "Name parameter is required")
	}
	return c.String(http.StatusOK, "Hello, "+name+"!")
}


func main () {
	e := echo.New()
	e.GET("/",WelcomeHandler)
	e.POST("/signup",SignupHandler)
	e.GET("/sayhi", SayHiHandler)
	e.Start(":8080")
}
