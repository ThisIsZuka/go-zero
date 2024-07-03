package userservicelogic

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"test.com/model"
	"test.com/rpc/internal/svc"
	"test.com/rpc/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNewUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNewUserLogic {
	return &AddNewUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddNewUserLogic) AddNewUser(in *user.AddNewUserRequest) (*user.AddNewUserResponse, error) {

	// แปลงค่า BDATE จาก string เป็น time.Time
	bdate, con_err := time.Parse("2006-01-02", in.User.Bdate) // สมมติว่า req.BDATE เป็นรูปแบบ "YYYY-MM-DD"
	if con_err != nil {
		l.Logger.Errorf("Failed to parse BDATE: %v", con_err)
		return nil, con_err
	}

	err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		ID:       primitive.NewObjectID(),
		FNAME:    in.User.Fname,
		LNAME:    in.User.Lname,
		AGE:      in.User.Age,
		BDATE:    bdate,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	// สร้างและส่ง AddNewUserResponse กลับ
	return &user.AddNewUserResponse{}, nil
}
