package falshsale

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashSaleOrderingGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashSaleOrderingGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashSaleOrderingGetLogic {
	return &FlashSaleOrderingGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashSaleOrderingGetLogic) FlashSaleOrderingGet(req *types.FlashSaleOrderingReq) (resp *types.FlashSaleOrderingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
