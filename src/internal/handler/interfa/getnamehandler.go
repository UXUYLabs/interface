package interfa

import (
	"net/http"

	"uxuy/src/internal/logic/interfa"
	"uxuy/src/internal/svc"
	"uxuy/src/internal/types"

	bizresponse "uxuy/src/util/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, bizresponse.ErrInvalidArgs.WithMessage(err.Error()))
			return
		}

		l := interfa.NewGetNameLogic(r.Context(), svcCtx)
		resp, err := l.GetName(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, bizresponse.NewSuccessResp(resp))
		}
	}
}
