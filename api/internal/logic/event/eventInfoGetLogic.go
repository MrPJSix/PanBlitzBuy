package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventInfoGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventInfoGetLogic {
	return &EventInfoGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventInfoGetLogic) EventInfoGet(req *types.EventInfoGetReq) (resp *types.EventInfoGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
