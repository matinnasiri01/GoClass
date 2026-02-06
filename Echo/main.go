package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ----------------------------
func Test(c echo.Context) error { 
	return c.String(http.StatusOK,"Test")
}
// ----------------------------
func WelcomeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}

// ----------------------------
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func SignupHandler(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

// ----------------------------
func SayHiHandler(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		return c.String(http.StatusBadRequest, "Name parameter is required")
	}
	return c.String(http.StatusOK, "Hello, "+name+"!")
}

// ----------------------------
func PathParameterHandler(c echo.Context) error {
	name := c.Param("namep")
	age := c.Param("agep")

	return c.String(
		http.StatusOK,
		fmt.Sprintf("Name : %s -----> Age -----> %s Grade -----> *****",
			name, age))
}

func main() {
	e := echo.New()
	// Base:
	e.GET("/", WelcomeHandler)
	e.POST("/signup", SignupHandler)
	e.GET("/sayhi", SayHiHandler)
	e.GET("/name/:namep/age/:agep", PathParameterHandler)

	// Routing:
	// "/stutent/*"
	// all urls start with `/stutent` return to the this rout handeler.
	// or Use group:
	g := e.Group("/users")
	g.GET("", Test)
	g.POST("", Test)
	// -> 8080:/users/:id
	g.GET("/:id", Test)

	/*
		Binding
		you can bind json in request body and use at 
		create an struck 
		create an object 
		use c.Bind(new_object)
		Done!
	
	*/

	e.Start(":8080")
}
