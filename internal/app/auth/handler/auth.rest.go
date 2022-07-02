package handler

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	errlib "github.com/nenecchuu/arcana/err"
	"github.com/nenecchuu/arcana/response"
	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/app/auth/model"
	tm "github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
	"github.com/nenecchuu/lizbeth-be-core/internal/util"
)

// Spotify Auth Callback
// @Summary Spotify Auth Callback
// @Description Spotify Auth Callback
// @Tags Auth
// @Accept json
// @Produce json
// @Param state query string true "state from callback"
// @Param user_id path string true "User ID"
// @Success 200 {object} response.JSONResponse{data=model.LinkageCallbackBodyRes}
// @Failure 400 {object} response.JSONResponse
// @Failure 500 {object} response.JSONResponse
// @Router /auth/callback [GET]
func (m *RestModule) HandleLinkageCallback(fc *fiber.Ctx) error {
	ctx, span := tracer.StartSpan(fc.Context(), "auth.rest.HandleLinkageCallback", nil)
	defer span.End()

	var (
		beginTs = time.Now()
		qParam  = &model.LinkageCallbackQParams{}
		res     response.JSONResponse
		err     error
		udata   *um.UserNoSqlSchema
		tdata   *tm.TokenNoSqlSchema
		resData *model.LinkageCallbackBodyRes
	)

	err = fc.QueryParser(qParam)

	if err != nil {
		e := errlib.GetError(err)
		return util.ReturnErrorToFiberResponse(fc, e)
	}

	udata, tdata, err = m.authUsecase.ProcessLinkageCallback(ctx, qParam.ToLinkageCallback())

	if err != nil {
		e := errlib.GetError(err)
		return util.ReturnErrorToFiberResponse(fc, e)
	}

	resData = model.BuildLinkageCallbackBodyRes(udata, tdata)

	res = response.NewJSONResponse().WithStatusCode(http.StatusOK).WithLatency(beginTs).WithData(resData)

	return fc.JSON(res)
}
