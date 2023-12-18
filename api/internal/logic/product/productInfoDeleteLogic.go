package product

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoDeleteLogic {
	return &ProductInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductInfoDeleteLogic) ProductInfoDelete(req *types.ProductInfoDeleteReq) (resp *types.ProductInfoDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
