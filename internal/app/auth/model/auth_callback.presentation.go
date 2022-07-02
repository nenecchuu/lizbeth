package model

import (
	tm "github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
)

type LinkageCallbackQParams struct {
	State string `json:"state"`
	Code  string `json:"code"`
}

type LinkageCallbackBodyRes struct {
	UserId        string `json:"_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	SpotifyUserId string `json:"spotify_user_id"`
	ImageUrl      string `json:"image_url"`
	AccessToken   string `json:"access_token"`
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
}

func BuildLinkageCallbackBodyRes(user *um.UserNoSqlSchema, token *tm.TokenNoSqlSchema) *LinkageCallbackBodyRes {
	return &LinkageCallbackBodyRes{
		UserId:        user.Id.Hex(),
		Name:          user.Name,
		Email:         user.Name,
		SpotifyUserId: user.SpotifyData.UserId,
		ImageUrl:      user.SpotifyData.ImageUrl,
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		ExpiresIn:     token.ExpiresIn,
	}
}

func (x *LinkageCallbackQParams) ToLinkageCallback() *LinkageCallback {
	return &LinkageCallback{
		Code:  x.Code,
		State: x.State,
	}
}
