package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditEventLogic {
	return &EditEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditEventLogic) EditEvent(in *productevent.EditEventReq) (*productevent.GeneralResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GeneralResp{}, nil
}
