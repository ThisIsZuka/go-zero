package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test.com/api/user/internal/logic/user"
	"test.com/api/user/internal/svc"
	"test.com/api/user/internal/types"
)

func DeleteUserByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteUserByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeleteUserByIdLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUserById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
