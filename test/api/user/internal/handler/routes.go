// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "test.com/api/user/internal/handler/user"
	"test.com/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/deleteUserById/:id",
				Handler: user.DeleteUserByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: user.UserHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getallusers",
				Handler: user.GetAllUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getusers",
				Handler: user.GetUserByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/newuser",
				Handler: user.AddNewUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}
