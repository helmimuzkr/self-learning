package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/helmimuzkr/templ/store/db"
	"github.com/helmimuzkr/templ/types"
	"github.com/helmimuzkr/templ/utill"
)

func UserInformationHandler(ctx *types.Context) error {
	name := ctx.Request.URL.Query().Get("name")
	if name == "" {
		name = "default"
	}

	user := types.User{
		Name: name,
		Age:  99,
	}

	data := types.Data{
		Status:  "found",
		Content: user,
	}

	return utill.TemplEngine.ExecuteTemplate(ctx.ResponseWriter, "index.html", data)
}

func UserListHandler(ctx *types.Context) error {
	user := types.User{
		Name: "jamal",
		Age:  123,
	}
	user2 := types.User{
		Name: "kanza",
		Age:  99,
	}
	users := []types.User{user, user2}

	data := types.Data{
		Status:  "found",
		Content: users,
	}

	return utill.TemplEngine.ExecuteTemplate(ctx.ResponseWriter, "list-users", data)
}

func RegisterUserHandler(ctx *types.Context) error {
	request := new(types.RegisterUserRequest)
	if err := utill.JsonDecode(ctx.Request, request); err != nil {
		utill.JsonErrorWriter(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
		return err
	}

	id := uuid.NewString()
	if err := db.SaveUser(request, id); err != nil {
		utill.JsonErrorWriter(ctx.ResponseWriter, http.StatusInternalServerError, err.Error())
		return err
	}

	return utill.JsonEncodeWriter(ctx.ResponseWriter, 200, "ok")
}
