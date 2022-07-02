package spotify_api

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/arcana/util"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	iutil "github.com/nenecchuu/lizbeth-be-core/internal/util"
	"github.com/rs/zerolog/log"
)

func (x *Module) GenerateToken(ctx context.Context, reqData *model.SpotifyAuthorizeData) (*model.SpotifyAuthorizeBodyRes, error) {
	_, span := tracer.StartSpan(ctx, "api_call.spotify.Authorize", nil)
	defer span.End()

	var (
		e            error
		reqBody      = reqData.ToSpotifyAuthorizeBodyReq(x.spotifyConfig.Credentials.AuthorizeCallbackUrl)
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
	hcRes, e = x.authHttpClient.PostForm(reqUrl, reqBody.ToMapString(), reqHeader)

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

func (x *Module) GenerateAuthorizeLink(ctx context.Context, user_id string) string {
	_, span := tracer.StartSpan(ctx, "api_call.spotify.GenerateAuthorizeLink", nil)
	defer span.End()

	var (
		params  = url.Values{}
		qParams string
		res     string
	)

	params.Add("response_type", "code")
	params.Add("client_id", x.spotifyConfig.Credentials.ClientId)
	params.Add("scope", x.spotifyConfig.Credentials.Scope)
	params.Add("state", user_id)
	params.Add("redirect_uri", x.spotifyConfig.Credentials.AuthorizeCallbackUrl)

	qParams = params.Encode()

	res = fmt.Sprintf("%s/%s?%s", x.spotifyConfig.AuthApi.HttpClient.BaseUrl, x.spotifyConfig.AuthApi.Endpoints.Authorize, qParams)

	return res
}
