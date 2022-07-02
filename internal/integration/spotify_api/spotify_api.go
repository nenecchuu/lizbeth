package spotify_api

import (
	"context"

	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
)

type SpotifyApiCallIntegration interface {
	GenerateToken(ctx context.Context, reqBody *model.SpotifyAuthorizeData) (*model.SpotifyAuthorizeBodyRes, error)
	GetUserInfo(ctx context.Context, accessToken string) (*model.SpotifyGetUserProfileBodyRes, error)
	GenerateAuthorizeLink(ctx context.Context, user_id string) string
}
