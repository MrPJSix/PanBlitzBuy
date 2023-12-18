package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEventLogic {
	return &DeleteEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteEventLogic) DeleteEvent(in *productevent.EventGeneralReq) (*productevent.GeneralResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GeneralResp{}, nil
}
