package usecase

import (
	"context"

	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type ChoreUsecase interface {
	ProcessInitConversation(ctx context.Context, ci cbm.ChatInfo) error
	ProcessWelcome(ctx context.Context, ci cbm.ChatInfo) error
}
