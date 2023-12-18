package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddEventLogic {
	return &AddEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddEventLogic) AddEvent(in *productevent.AddEventReq) (*productevent.GeneralResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GeneralResp{}, nil
}
