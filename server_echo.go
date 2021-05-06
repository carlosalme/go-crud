package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "fmt"
)



func main() {
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World!")

  })
  e.Logger.Fatal(e.Start(":1323"))

  e.POST("/save", func(c echo.Context) error {
   name := c.FormValue("name")
   email := c.FormValue("email")
   return c.String(http.StatusOK, "name:" + name + ", email:" + email)
   })

  e.GET("/users/:id", func(c echo.Context) error {
   id := c.Param("id")
   return c.String(http.StatusOK, id)
   })

  //e.PUT("/users/:id", updateUser)
  //e.DELETE("/users/:id", deleteUser)
}

// e.GET("/show", show)
func show(c echo.Context) error {
 // Get team and member from the query string 
 team := c.QueryParam("team")
 member := c.QueryParam("member")
 return c.String(http.StatusOK, "team:" + team + ", member: " + member)
}

