package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"html/template"
	"io"
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



// ----------------------------
type TemplateRender struct {
    templates *template.Template
}
func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}


type Author struct {
    FirstName string
    LastName  string
    Age       int
}
func aboutHandler(c echo.Context) error {
    jkRowling := Author{
        FirstName: "Joanne",
        LastName:  "Rowling",
        Age:       58,
    }

    data := map[string]interface{}{
        "author": jkRowling,
        "books": []string{"Harry Potter and the Sorcerer's Stone", "Harry Potter and the Chamber of Secrets",
            "Harry Potter and the Prisoner of Azkaban", "Harry Potter and the Goblet of Fire", "Harry Potter and the Order of the Phoenix",
            "Harry Potter and the Half-Blood Prince", "Harry Potter and the Deathly Hallows",
        },
        "isFamous": true,
    }
    return c.Render(http.StatusOK, "about.html", data)
}


func main() {
	e := echo.New()
	e.Renderer = &TemplateRender{
     	templates: template.Must(template.ParseGlob("templates/*.html")),
     }
	e.GET("/about", aboutHandler)

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
