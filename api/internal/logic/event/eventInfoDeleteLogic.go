package event

import (
	"context"

	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EventInfoDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEventInfoDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventInfoDeleteLogic {
	return &EventInfoDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventInfoDeleteLogic) EventInfoDelete(req *types.EventInfoDeleteReq) (resp *types.EventInfoDeleteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
