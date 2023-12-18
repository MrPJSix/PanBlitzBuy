package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventListGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventListGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventListGetLogic {
	return &EventListGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventListGetLogic) EventListGet(req *types.EventListGetReq) (resp *types.EventListGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
