package user

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserEmailSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserEmailSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEmailSendLogic {
	return &UserEmailSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserEmailSendLogic) UserEmailSend(req *types.UserEmailSendReq) (resp *types.UserEmailSendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
