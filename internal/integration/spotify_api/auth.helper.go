package spotify_api

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/arcana/util"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	iutil "github.com/nenecchuu/lizbeth-be-core/internal/util"
	"github.com/rs/zerolog/log"
)

func (x *Module) requestToken(ctx context.Context, reqBody map[string]string) (*model.SpotifyAuthorizeBodyRes, error) {
	_, span := tracer.StartSpan(ctx, "api_call.spotify.requestToken", nil)
	defer span.End()

	var (
		e            error
		reqHeader    = map[string]string{}
		reqUrl       = fmt.Sprintf("%s/%s", x.spotifyConfig.AuthApi.HttpClient.BaseUrl, x.spotifyConfig.AuthApi.Endpoints.Token)
		hcRes        *http.Response
		jsonRes      = &model.SpotifyAuthorizeBodyRes{}
		token        = fmt.Sprintf("%s:%s", x.spotifyConfig.Credentials.ClientId, x.spotifyConfig.Credentials.ClientSecret)
		encodedToken string
	)

	encodedToken = base64.StdEncoding.EncodeToString([]byte(token))
	reqHeader["Authorization"] = fmt.Sprintf("Basic %s", encodedToken)
	reqHeader["Content-Type"] = "application/x-www-form-urlencoded"
	hcRes, e = x.authHttpClient.PostForm(reqUrl, reqBody, reqHeader)

	if e != nil {
		log.Err(e).Interface("body", reqBody).Interface("headers", reqHeader).Msg(e.Error())
		return nil, e
	}

	if hcRes.StatusCode > 299 {
		e = iutil.NewBadRequestErr("Request failed")
		return nil, e
	}

	e = util.ParseResponseBodyToJson(hcRes, jsonRes)

	if e != nil {
		log.Err(e).Interface("body", reqBody).Interface("headers", reqHeader).Msg(e.Error())
		return nil, e
	}

	return jsonRes, nil
}
