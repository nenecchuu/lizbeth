package handler

import (
	"github.com/nenecchuu/lizbeth-be-core/internal/app/player/usecase"
)

type ChatbotModule struct {
	playerUsecase usecase.PlayerUsecase
}

type ChatbotOpts struct {
	PlayerUsecase usecase.PlayerUsecase
}

func NewChatbot(o ChatbotOpts) *ChatbotModule {
	return &ChatbotModule{
		playerUsecase: o.PlayerUsecase,
	}
}
