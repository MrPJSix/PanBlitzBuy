package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFrontEventLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFrontEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFrontEventLogic {
	return &GetFrontEventLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFrontEventLogic) GetFrontEvent(in *productevent.EventGeneralReq) (*productevent.GetFrontEventResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GetFrontEventResp{}, nil
}
