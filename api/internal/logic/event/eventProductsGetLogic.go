package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventProductsGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventProductsGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventProductsGetLogic {
	return &EventProductsGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventProductsGetLogic) EventProductsGet() (resp *types.EventProductsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
