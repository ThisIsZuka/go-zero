package userservicelogic

import (
	"context"

	"test.com/rpc/internal/svc"
	"test.com/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetUserByIdResponse{
		Id:       res.ID.Hex(),
		Fname:    res.FNAME,
		Lname:    res.LNAME,
		Age:      res.AGE,
		Bdate:    res.BDATE.String(),
		CreateAt: res.CreateAt.String(),
		UpdateAt: res.UpdateAt.String(),
	}, nil
}
