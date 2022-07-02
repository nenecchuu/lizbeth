package repository

import (
	"context"

	"github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
)

type TokenRepository interface {
	StoreToken(ctx context.Context, token *model.TokenNoSqlSchema) error
}
