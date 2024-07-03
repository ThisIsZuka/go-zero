package user

import (
	"context"

	"test.com/api/user/internal/svc"
	"test.com/api/user/internal/types"
	"test.com/rpc/client/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNewUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddNewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNewUserLogic {
	return &AddNewUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddNewUserLogic) AddNewUser(req *types.NewUserRequest) (resp *types.NewUserResponse, err error) {
	rpcReq := &userservice.AddNewUserRequest{
		User: &userservice.User{
			Fname: req.FNAME,
			Lname: req.LNAME,
			Age:   req.AGE,
			Bdate: req.BDATE,
		},
	}

	_, err = l.svcCtx.UserClient.AddNewUser(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &types.NewUserResponse{
		Message: "Successful",
	}, nil
}
