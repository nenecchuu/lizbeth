package model

type SpotifyAuthorizeData struct {
	Code string
}

func (x *SpotifyAuthorizeData) ToSpotifyAuthorizeBodyReq(redirectUri string) *SpotifyAuthorizeBodyReq {
	return &SpotifyAuthorizeBodyReq{
		RedirectUri: "http://localhost:8000/auth/callback",
		GrantType:   "authorization_code",
		Code:        x.Code,
	}
}
