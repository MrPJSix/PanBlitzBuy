package falshsale

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashSaleResultGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashSaleResultGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashSaleResultGetLogic {
	return &FlashSaleResultGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashSaleResultGetLogic) FlashSaleResultGet() (resp *types.FlashSaleResultResp, err error) {
	// todo: add your logic here and delete this line

	return
}
