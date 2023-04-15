package chatbot

import (
	sr "github.com/nenecchuu/lizbeth-be-core/internal/app/session/repository"
	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type ChatbotMiddlewareModule struct {
	sessionRepository sr.SessionRepository
	userRepository    ur.UserRepository
	tokenRepository   tr.TokenRepository
	chatbotManager    telegram_bot.TelegramBotIntegration
}

type ChatbotMiddlewareOpts struct {
	TokenRepository   tr.TokenRepository
	UserRepository    ur.UserRepository
	SessionRepository sr.SessionRepository
	ChatbotManager    telegram_bot.TelegramBotIntegration
}

func NewChatbotMiddleware(o ChatbotMiddlewareOpts) *ChatbotMiddlewareModule {
	return &ChatbotMiddlewareModule{
		sessionRepository: o.SessionRepository,
		tokenRepository:   o.TokenRepository,
		userRepository:    o.UserRepository,
		chatbotManager:    o.ChatbotManager,
	}
}
