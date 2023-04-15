package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	"github.com/nenecchuu/lizbeth-be-core/internal/middleware/chatbot"
)

func (i *Initiator) InitChatbotMiddleware(cfg *config.MainConfig, infra *service.Infrastructure, integration *service.Integration, repos *service.Repositories) *service.Middlewares {
	cbmw := chatbot.NewChatbotMiddleware(chatbot.ChatbotMiddlewareOpts{
		TokenRepository:   repos.TokenRepository,
		UserRepository:    repos.UserRepository,
		SessionRepository: repos.SessionRepository,
		ChatbotManager:    integration.TelegramBotManager,
	})

	return &service.Middlewares{
		ChatbotMiddleware: cbmw,
	}
}
