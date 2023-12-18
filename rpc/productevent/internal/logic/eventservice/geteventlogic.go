package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEventLogic {
	return &GetEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEventLogic) GetEvent(in *productevent.EventGeneralReq) (*productevent.GetEventResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GetEventResp{}, nil
}
