package spotify_api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/arcana/util"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/constants"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	iutil "github.com/nenecchuu/lizbeth-be-core/internal/util"
	"github.com/rs/zerolog/log"
)

func (x *Module) searchForItem(ctx context.Context, accessToken string, keyword string, searchType constants.SpotifySearchType) (*model.SpotifySearchResultBodyRes, error) {
	_, span := tracer.StartSpan(ctx, "api_call.spotify.SearchForItem", nil)
	defer span.End()

	var (
		e         error
		qParams   = url.Values{}
		reqHeader = map[string]string{}
		reqUrl    = fmt.Sprintf("%s/%s", x.spotifyConfig.CoreApi.HttpClient.BaseUrl, x.spotifyConfig.CoreApi.Endpoints.Search)
		hcRes     *http.Response
		jsonRes   = &model.SpotifySearchResultBodyRes{}
	)

	reqHeader["Authorization"] = fmt.Sprintf("Bearer %s", accessToken)

	qParams.Add("q", keyword)
	qParams.Add("searchType", string(searchType))

	hcRes, e = x.coreHttpClient.Get(reqUrl, reqHeader)

	if e != nil {
		log.Err(e).Str("qParams", qParams.Encode()).Interface("headers", reqHeader).Msg(e.Error())
		return nil, e
	}

	if hcRes.StatusCode > 299 {
		e = iutil.NewBadRequestErr("Request failed")
		return nil, e
	}

	e = util.ParseResponseBodyToJson(hcRes, jsonRes)

	if e != nil {
		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return jsonRes, nil
}
