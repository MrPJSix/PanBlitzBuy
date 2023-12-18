package product

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"pan-blitz-buy/api/internal/logic/product"
	"pan-blitz-buy/api/internal/svc"
	"pan-blitz-buy/api/internal/types"
)

func ProductInfoUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductInfoUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := product.NewProductInfoUpdateLogic(r.Context(), svcCtx)
		resp, err := l.ProductInfoUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
