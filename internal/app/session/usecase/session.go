package usecase

import (
	"context"

	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type SessionUsecase interface {
	ProcessCreateNewSession(ctx context.Context, ci gm.ChatInfo) error
}
