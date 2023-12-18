package product

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductInfoAddLogic {
	return &ProductInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductInfoAddLogic) ProductInfoAdd(req *types.ProductInfoAddReq) (resp *types.ProductInfoAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
