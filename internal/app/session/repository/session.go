package repository

import (
	"context"

	"github.com/nenecchuu/lizbeth-be-core/internal/app/session/model"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SessionRepository interface {
	FindSessionByChatbotSessionId(ctx context.Context, chatbotSessionId string, chatbotChannel constants.ChatbotChannelEnum) (*model.SessionNoSqlSchema, error)
	FindSessionById(ctx context.Context, id primitive.ObjectID) (*model.SessionNoSqlSchema, error)
	UpdateSession(ctx context.Context, id primitive.ObjectID, session *model.SessionNoSqlSchema) error
	StoreSession(ctx context.Context, session *model.SessionNoSqlSchema) error
}
