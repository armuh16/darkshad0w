package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	//your solution here
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	users[id].Email = u.Email
	users[id].Password = u.Password
	return c.JSON(http.StatusOK, users[id])
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid ID")
	}
	if id < len(users) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"massage":  "succes get user by id",
			"id":       users[id].Id,
			"name":     users[id].Name,
			"email":    users[id].Email,
			"password": users[id].Password,
		})

	}
	return c.String(http.StatusNotFound, "Not Found")
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id",
		})
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			if i == len(users)-1 {
				users = users[:len(users)-1]
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success get all users",
					"users":    users,
				})
			}
			users = users[i+1:]
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get all users",
				"users":    users,
			})
		}

	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "invalid id",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/user/:id", GetUserController)
	e.GET("/users", GetUsersController)
	e.DELETE("/user/:id", DeleteUserController)
	e.POST("/user", CreateUserController)
	e.PUT("/user/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
