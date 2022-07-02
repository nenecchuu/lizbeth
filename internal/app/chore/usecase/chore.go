package usecase

import (
	"context"

	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type ChoreUsecase interface {
	ProcessInitConversation(ctx context.Context, ci gm.ChatInfo) error
	ProcessWelcome(ctx context.Context, ci gm.ChatInfo) error
}
