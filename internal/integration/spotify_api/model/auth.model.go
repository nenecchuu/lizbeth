package model

import "github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/constants"

type SpotifyAuthorizeData struct {
	Code string
}

type SpotifyRefreshTokenData struct {
	RefreshToken string
}

func (x *SpotifyAuthorizeData) ToSpotifyAuthorizeBodyReq(redirectUri string) *SpotifyAuthorizeBodyReq {
	return &SpotifyAuthorizeBodyReq{
		RedirectUri: redirectUri,
		GrantType:   string(constants.SpotifyGrantTypeAuthorizationCode),
		Code:        x.Code,
	}
}

func (x *SpotifyRefreshTokenData) ToSpotifyAuthorizeBodyReq() *SpotifyRefreshTokenBodyReq {
	return &SpotifyRefreshTokenBodyReq{
		RefreshToken: x.RefreshToken,
		GrantType:    string(constants.SpotifyGrantTypeRefreshToken),
	}
}
