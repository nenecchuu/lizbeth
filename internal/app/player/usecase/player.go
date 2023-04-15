package usecase

import (
	"context"

	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type PlayerUsecase interface {
	ProcessQueueTrack(context.Context, cbm.ChatInfo) error
}
