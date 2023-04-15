package repository

import (
	"context"

	"github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenRepository interface {
	FindAndValidateTokenByUserId(ctx context.Context, id primitive.ObjectID) (*model.TokenNoSqlSchema, error)
	StoreToken(ctx context.Context, token *model.TokenNoSqlSchema) error
}
