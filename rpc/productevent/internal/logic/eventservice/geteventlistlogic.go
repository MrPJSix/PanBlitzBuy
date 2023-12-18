package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEventListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEventListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEventListLogic {
	return &GetEventListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEventListLogic) GetEventList(in *productevent.GetEventListReq) (*productevent.GetEventListResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GetEventListResp{}, nil
}
