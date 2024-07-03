package userservicelogic

import (
	"context"
	"strconv"

	"test.com/rpc/internal/svc"
	"test.com/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserByIdLogic {
	return &DeleteUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserByIdLogic) DeleteUserById(in *user.DeleteUserByIdRequest) (*user.DeleteUserByIdResponse, error) {

	rpcRes, err := l.svcCtx.UserModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.DeleteUserByIdResponse{
		Message: strconv.FormatInt(rpcRes, 10),
	}, nil
}
