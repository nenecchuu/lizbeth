package usecase

import (
	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type Module struct {
	tokenRepository tr.TokenRepository
	userRepository  ur.UserRepository
	chatbotManager  telegram_bot.TelegramBotIntegration
}

type Opts struct {
	TokenRepository tr.TokenRepository
	UserRepository  ur.UserRepository
	ChatbotManager  telegram_bot.TelegramBotIntegration
}

func New(o Opts) *Module {
	return &Module{
		tokenRepository: o.TokenRepository,
		userRepository:  o.UserRepository,
		chatbotManager:  o.ChatbotManager,
	}
}
