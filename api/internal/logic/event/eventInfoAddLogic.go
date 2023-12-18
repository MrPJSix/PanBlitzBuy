package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventInfoAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventInfoAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventInfoAddLogic {
	return &EventInfoAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventInfoAddLogic) EventInfoAdd(req *types.EventInfoAddReq) (resp *types.EventInfoAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
