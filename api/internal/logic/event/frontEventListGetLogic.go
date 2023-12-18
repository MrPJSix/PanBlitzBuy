package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FrontEventListGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFrontEventListGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FrontEventListGetLogic {
	return &FrontEventListGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FrontEventListGetLogic) FrontEventListGet(req *types.FrontEventListGetReq) (resp *types.FrontEventListGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
