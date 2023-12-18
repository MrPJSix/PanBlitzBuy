package event

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pan-blitz-buy/api/internal/logic/event"
	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"
)

func FrontEventInfoGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FrontEventInfoGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := event.NewFrontEventInfoGetLogic(r.Context(), svcCtx)
		resp, err := l.FrontEventInfoGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
