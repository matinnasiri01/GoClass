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



func main () {
	e := echo.New()
	e.GET("/",WelcomeHandler)
	e.POST("/signup",SignupHandler)
	e.Start(":8080")
}
