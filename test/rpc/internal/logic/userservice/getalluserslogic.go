package userservicelogic

import (
	"context"

	"test.com/rpc/internal/svc"
	"test.com/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllUsersLogic {
	return &GetAllUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllUsersLogic) GetAllUsers(in *user.GetAllUserRequest) (*user.GetAllUsersResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}

	var userResponse []*user.User
	for _, usr := range res {
		userResponse = append(userResponse, &user.User{
			Id:       usr.ID.Hex(),
			Fname:    usr.FNAME,
			Lname:    usr.LNAME,
			Age:      usr.AGE,
			Bdate:    usr.BDATE.String(),
			CreateAt: usr.CreateAt.String(),
			UpdateAt: usr.UpdateAt.String(),
		})
	}

	return &user.GetAllUsersResponse{
		Users: userResponse,
	}, nil
}
