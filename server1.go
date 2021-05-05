package main

import (
  "net/http"
  //"html/template"
  //"log"
  "github.com/labstack/echo/v4"
)

//var tmpl *template.Template

func main() {
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World!")

  })
  e.Logger.Fatal(e.Start(":1323"))
  e.POST("/users", saveUser)
  e.GET("/users/:id", getUser)
  //e.PUT("/users/:id", updateUser)
  //e.DELETE("/users/:id", deleteUser)
   
/*
  fs := http.FileServer(http.Dir("."))
  http.Handle("/", fs)
  log.Println("listening on :3000...")
  err := http.ListenAndServe(":3000", nil)
  if err != nil {
    log.Fatal(err)
  }
  */
}

//e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
 // User ID from path 'users/:id
 id := c.Param("id")
 return c.String(http.StatusOK, id)
}

// e.GET("/show", show)
func show(c echo.Context) error {
 // Get team and member from the query string 
 team := c.QueryParam("team")
 member := c.QueryParam("member")
 return c.String(http.StatusOK, "team:" + team + ", member: " + member)
}

// e.POST("/save", saveUser)
func saveUser(c echo.Context) error {
 //Get name and email
 name := c.FormValue("name")
 email := c.FormValue("email")
 return c.String(http.StatusOK, "name: " + name + ", email:" + email)
}


