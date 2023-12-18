package productservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSelectedProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSelectedProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSelectedProductListLogic {
	return &GetSelectedProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSelectedProductListLogic) GetSelectedProductList(in *productevent.ProductGeneralReq) (*productevent.GetProductListResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GetProductListResp{}, nil
}
