package falshsale

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pan-blitz-buy/api/internal/logic/falshsale"
	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"
)

func FlashSaleOrderingGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FlashSaleOrderingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := falshsale.NewFlashSaleOrderingGetLogic(r.Context(), svcCtx)
		resp, err := l.FlashSaleOrderingGet(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
