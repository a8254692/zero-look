package userbehavior

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"minicode.com/sirius/go-back-server/service/userbehavior/api/internal/logic/userbehavior"
	"minicode.com/sirius/go-back-server/service/userbehavior/api/internal/svc"
	"minicode.com/sirius/go-back-server/service/userbehavior/api/internal/types"

	"minicode.com/sirius/go-back-server/utils/response"
)

func AddShareNumHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddShareNumReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := userbehavior.NewAddShareNumLogic(r.Context(), svcCtx)
		resp, err := l.AddShareNum(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.SuccessResponse(w, resp)
		}
	}
}
