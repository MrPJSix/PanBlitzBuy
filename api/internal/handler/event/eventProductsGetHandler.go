package event

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pan-blitz-buy/api/internal/logic/event"
	"pan-blitz-buy/api/internal/svc"
)

func EventProductsGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := event.NewEventProductsGetLogic(r.Context(), svcCtx)
		resp, err := l.EventProductsGet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
