package usecase

import (
	"context"

	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type SessionUsecase interface {
	ProcessCreateNewSession(ctx context.Context, ci cbm.ChatInfo) error
}
