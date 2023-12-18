package product

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoGetLogic {
	return &ProductInfoGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductInfoGetLogic) ProductInfoGet(req *types.ProductInfoGetReq) (resp *types.ProductInfoGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
