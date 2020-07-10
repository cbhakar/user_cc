package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/user/profile",UserProfileHandler)
	e.GET("/microservice/name",MicroserviceNameHandler)

	e.Logger.Fatal(e.Start(":9992"))
}



// using hard coded user details
type User struct {
	Name string `json:"name"`
	DOB  string	`json:"dob"`
	Age  int	`json:"age"`
	PhoneNumber 	string `json:"phone-number"`
}


func MicroserviceNameHandler(c echo.Context) error {
	return c.String(http.StatusOK, "welcome to user microservice")
}


func UserProfileHandler(c echo.Context) error {
	uName := c.QueryParam("q")
	if len(uName) < 1 {
		return c.JSON(http.StatusBadRequest, "incorrect credentials")
	}

	u :=User{
		Name:        uName,
		DOB:         "12-01-1999",
		Age:         24,
		PhoneNumber: "+919999999999",
	}

	return c.JSON(http.StatusOK, u)
}