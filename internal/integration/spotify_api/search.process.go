package spotify_api

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/constants"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	"github.com/rs/zerolog/log"
)

func (x *Module) SearchTracks(ctx context.Context, accessToken string, keyword string) (*model.SpotifyTrackSearchBodyRes, error) {
	_, span := tracer.StartSpan(ctx, "api_call.spotify.SearchTracks", nil)
	defer span.End()

	var (
		e   error
		res = &model.SpotifySearchResultBodyRes{}
	)

	res, e = x.searchForItem(ctx, accessToken, keyword, constants.SpotifySearchTypeTrack)

	if e != nil {

		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return &res.Tracks, nil
}
