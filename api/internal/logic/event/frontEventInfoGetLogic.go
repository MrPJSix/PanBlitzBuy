package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FrontEventInfoGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFrontEventInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FrontEventInfoGetLogic {
	return &FrontEventInfoGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FrontEventInfoGetLogic) FrontEventInfoGet(req *types.FrontEventInfoGetReq) (resp *types.FrontEventInfoGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
