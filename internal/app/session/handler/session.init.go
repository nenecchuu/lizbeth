package handler

import (
	"github.com/nenecchuu/lizbeth-be-core/internal/app/auth/usecase"
)

type ChatbotModule struct {
	authUsecase usecase.AuthUsecase
}

type ChatbotOpts struct {
	AuthUsecase usecase.AuthUsecase
}

func NewChatbot(o ChatbotOpts) *ChatbotModule {
	return &ChatbotModule{
		authUsecase: o.AuthUsecase,
	}
}
