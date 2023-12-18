package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventInfoUpdateLogic {
	return &EventInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventInfoUpdateLogic) EventInfoUpdate(req *types.EventInfoUpdateReq) (resp *types.EventInfoUpdateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
