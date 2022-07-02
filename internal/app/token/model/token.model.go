package model

import (
	sam "github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenNoSqlSchema struct {
	UserId       primitive.ObjectID `bson:"user_id"`
	AccessToken  string             `bson:"access_token"`
	RefreshToken string             `bson:"refresh_token"`
	ExpiresIn    int                `bson:"expires_in"`
	TokenType    string             `bson:"token_type"`
	Scope        string             `bson:"scope"`
	Issuer       string             `bson:"issuer"`
}

func (x *TokenNoSqlSchema) BuildFromSpotifyAuthorizeBodyRes(spr *sam.SpotifyAuthorizeBodyRes, userId primitive.ObjectID) *TokenNoSqlSchema {
	return &TokenNoSqlSchema{
		UserId:       userId,
		AccessToken:  spr.AccessToken,
		RefreshToken: spr.RefreshToken,
		ExpiresIn:    spr.ExpiresIn,
		TokenType:    spr.TokenType,
		Scope:        spr.Scope,
		Issuer:       "spotify",
	}
}
