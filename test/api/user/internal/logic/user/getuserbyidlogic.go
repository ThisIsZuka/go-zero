package user

import (
	"context"

	"test.com/api/user/internal/svc"
	"test.com/api/user/internal/types"
	"test.com/rpc/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdRequest) (resp *types.GetUserByIdResponse, err error) {
	// สร้างคำขอ RPC request
	rpcReq := &userservice.GetUserByIdRequest{
		Id: req.ID,
	}

	// เรียกใช้ RPC service
	rpcResp, err := l.svcCtx.UserClient.GetUserById(l.ctx, rpcReq)
	if err != nil {
		l.Logger.Errorf("Failed to call GetUserById: %v", err)
		return nil, err
	}

	// สร้าง response จาก RPC response
	resp = &types.GetUserByIdResponse{
		ID:       rpcResp.Id,
		FNAME:    rpcResp.Fname,
		LNAME:    rpcResp.Lname,
		AGE:      rpcResp.Age,
		BDATE:    rpcResp.Bdate,
		CreateAt: rpcResp.CreateAt,
		UpdateAt: rpcResp.UpdateAt,
	}

	return resp, nil
}
