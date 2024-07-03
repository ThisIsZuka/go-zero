package user

import (
	"context"

	"test.com/api/user/internal/svc"
	"test.com/api/user/internal/types"
	"test.com/rpc/client/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserByIdLogic {
	return &DeleteUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserByIdLogic) DeleteUserById(req *types.DeleteUserByIdRequest) (resp *types.DeleteUserByIdResponse, err error) {
	rpcReq := &userservice.DeleteUserByIdRequest{
		Id: req.ID,
	}

	rpcRes, err := l.svcCtx.UserClient.DeleteUserById(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return &types.DeleteUserByIdResponse{
		Message: rpcRes.Message,
	}, nil
}
