package model

import (
	"time"

	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	sam "github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api/model"
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserNoSqlSchema struct {
	Id              primitive.ObjectID           `bson:"_id"`
	Name            string                       `bson:"name"`
	ChatId          string                       `bson:"chat_id"`
	ChatbotUserId   string                       `bson:"chatbot_user_id"`
	ChatbotChannel  constants.ChatbotChannelEnum `bson:"chatbot_channel"`
	ActiveSessionId primitive.ObjectID           `bson:"active_session_id"`
	SpotifyData     UserSpotifyDataNoSqlSchema   `bson:"spotify_data"`
}

type UserSpotifyDataNoSqlSchema struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	UserId   string `bson:"user_id"`
	ImageUrl string `bson:"image_url"`
	Product  string `bson:"product"`
	UserType string `bson:"user_type"`
}

func (x *UserNoSqlSchema) AssignSpotifyData(spr *sam.SpotifyGetUserProfileBodyRes) {
	x.SpotifyData = UserSpotifyDataNoSqlSchema{
		Name:     spr.DisplayName,
		Email:    spr.Email,
		UserId:   spr.Id,
		ImageUrl: spr.Images[0].Url,
		Product:  spr.Product,
		UserType: spr.Type,
	}
}

func (x *UserNoSqlSchema) BuildFromChatInfo(ci gm.ChatInfo) *UserNoSqlSchema {
	return &UserNoSqlSchema{
		Id:             primitive.NewObjectIDFromTimestamp(time.Now()),
		ChatId:         ci.ChatId,
		ChatbotUserId:  ci.SenderId,
		Name:           ci.SenderFullName,
		ChatbotChannel: ci.Channel,
	}
}
