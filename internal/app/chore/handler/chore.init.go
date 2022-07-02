package handler

import (
	"github.com/nenecchuu/lizbeth-be-core/internal/app/chore/usecase"
)

type ChatbotModule struct {
	choreUsecase usecase.ChoreUsecase
}

type ChatbotOpts struct {
	ChoreUsecase usecase.ChoreUsecase
}

func NewChatbot(o ChatbotOpts) *ChatbotModule {
	return &ChatbotModule{
		choreUsecase: o.ChoreUsecase,
	}
}
