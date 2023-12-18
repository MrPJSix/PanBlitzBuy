package falshsale

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pan-blitz-buy/api/internal/logic/falshsale"
	"pan-blitz-buy/api/internal/svc"
)

func FlashSaleResultGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := falshsale.NewFlashSaleResultGetLogic(r.Context(), svcCtx)
		resp, err := l.FlashSaleResultGet()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
