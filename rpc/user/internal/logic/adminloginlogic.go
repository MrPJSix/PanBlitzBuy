package logic

import (
	"context"

	"pan-blitz-buy/rpc/user/internal/svc"
	"pan-blitz-buy/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginLogic) AdminLogin(in *user.GeneralReq) (*user.GeneralResp, error) {
	// todo: add your logic here and delete this line

	return &user.GeneralResp{}, nil
}
