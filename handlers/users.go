package handlers

import (
	"net/http"

	"github.com/jpw547/indecision/structs"
	"github.com/labstack/echo"
)

// CreateUser adds a user to the database
func CreateUser(ctx echo.Context) error {
	// get the type from the context
	var u structs.User
	err := ctx.Bind(&u)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}
	err = store.CreateUser(&u)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(http.StatusOK, "OK")
}

// GetUser returns the specific user by their username
func GetUser(ctx echo.Context) error {
	// get the type from the context
	username := ctx.Param("username")

	user := structs.User{
		ID:       "",
		Username: username,
	}
	val, err := store.GetUser(&user)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, val)
}

// DeleteUserByUsername deletes the user with the given username
func DeleteUserByUsername(ctx echo.Context) error {
	// get the type from the context
	username := ctx.Param("username")
	err := store.DeleteUserByUsername(username)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "OK")
}
