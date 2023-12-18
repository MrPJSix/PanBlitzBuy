package logic

import (
	"context"

	"pan-blitz-buy/rpc/flashsale/flashsale"
	"pan-blitz-buy/rpc/flashsale/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FrontFlashsaleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFrontFlashsaleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FrontFlashsaleLogic {
	return &FrontFlashsaleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FrontFlashsaleLogic) FrontFlashsale(in *flashsale.FlashsaleReq) (*flashsale.GeneralResp, error) {
	// todo: add your logic here and delete this line

	return &flashsale.GeneralResp{}, nil
}
