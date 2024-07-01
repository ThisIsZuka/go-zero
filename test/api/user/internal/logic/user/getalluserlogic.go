package user

import (
	"context"
	"fmt"

	"test.com/api/user/internal/svc"
	"test.com/api/user/internal/types"
	"test.com/rpc/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUserLogic {
	return &GetAllUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllUserLogic) GetAllUser(req *types.GetAllUserRequest) (resp *types.GetAllUsersResponse, err error) {
	rpcReq := &userservice.GetAllUserRequest{}

	rpcResp, err := l.svcCtx.UserClient.GetAllUsers(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("Failed to call GetAllUsers: %v", err)
		return nil, err
	}
	fmt.Printf("%v", rpcResp.Users)
	// สร้าง response จาก RPC response
	var users []types.User
	for _, u := range rpcResp.Users {
		users = append(users, types.User{
			ID:       u.Id,
			FNAME:    u.Fname,
			LNAME:    u.Lname,
			AGE:      u.Age,
			BDATE:    u.Bdate,
			CreateAt: u.CreateAt,
			UpdateAt: u.UpdateAt,
		})
	}

	return &types.GetAllUsersResponse{
		Users: users,
	}, nil

}
