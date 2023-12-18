package admin

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetUsersLogic {
	return &AdminGetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetUsersLogic) AdminGetUsers(req *types.AdminUsersGetReq) (resp *types.AdminUsersGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
