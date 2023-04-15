package usecase

import (
	"context"

	"github.com/nenecchuu/lizbeth-be-core/internal/app/auth/model"
	tm "github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type AuthUsecase interface {
	ProcessHostAuthentication(ctx context.Context, ci cbm.ChatInfo) error
	ProcessLinkageCallback(ctx context.Context, data *model.LinkageCallback) (*um.UserNoSqlSchema, *tm.TokenNoSqlSchema, error)
}
