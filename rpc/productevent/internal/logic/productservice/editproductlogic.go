package productservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditProductLogic {
	return &EditProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditProductLogic) EditProduct(in *productevent.EditProductReq) (*productevent.GeneralResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GeneralResp{}, nil
}
