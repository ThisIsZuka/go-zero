package logic

import (
	"context"
	"fmt"

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
	fmt.Println("Received ID in API Layer:", in.Id)
	usr, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &user.GetUserByIdResponse{
		Id:    usr.ID.Hex(),
		Fname: usr.FNAME,
		Lname: usr.LNAME,
		Age:   usr.AGE,
		// Bdate:    usr.BDATE.Format(time.RFC3339),    // Convert time.Time to string
		// CreateAt: usr.CreateAt.Format(time.RFC3339), // Convert time.Time to string
		// UpdateAt: usr.UpdateAt.Format(time.RFC3339), // Convert time.Time to string
	}, nil
}
