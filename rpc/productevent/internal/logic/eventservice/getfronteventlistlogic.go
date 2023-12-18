package eventservicelogic

import (
	"context"

	"pan-blitz-buy/rpc/productevent/internal/svc"
	"pan-blitz-buy/rpc/productevent/productevent"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFrontEventListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFrontEventListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFrontEventListLogic {
	return &GetFrontEventListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFrontEventListLogic) GetFrontEventList(in *productevent.GetFrontEventListReq) (*productevent.GetFrontEventListResp, error) {
	// todo: add your logic here and delete this line

	return &productevent.GetFrontEventListResp{}, nil
}
