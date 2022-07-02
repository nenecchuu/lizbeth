package repository

import (
	"context"

	"github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	FindUserByChatbotUserId(ctx context.Context, chatbotUserId string, chatbotChannel constants.ChatbotChannelEnum) (*model.UserNoSqlSchema, error)
	FindUserById(ctx context.Context, id primitive.ObjectID) (*model.UserNoSqlSchema, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, user *model.UserNoSqlSchema) error
	StoreUser(ctx context.Context, user *model.UserNoSqlSchema) error
}
