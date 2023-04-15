package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nenecchuu/arcana/mongo"
	"github.com/nenecchuu/lizbeth-be-core/config"

	ah "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/handler"
	auc "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/usecase"
	"github.com/nenecchuu/lizbeth-be-core/internal/middleware/chatbot"

	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"

	chh "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/handler"
	chuc "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/usecase"

	sh "github.com/nenecchuu/lizbeth-be-core/internal/app/session/handler"
	sr "github.com/nenecchuu/lizbeth-be-core/internal/app/session/repository"
	suc "github.com/nenecchuu/lizbeth-be-core/internal/app/session/usecase"

	ph "github.com/nenecchuu/lizbeth-be-core/internal/app/player/handler"
	puc "github.com/nenecchuu/lizbeth-be-core/internal/app/player/usecase"

	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type Infrastructure struct {
	ChatbotModule *tgbotapi.BotAPI
	Mongo         *mongo.Database
}

type Repositories struct {
	TokenRepository   tr.TokenRepository
	UserRepository    ur.UserRepository
	SessionRepository sr.SessionRepository
}

type Usecases struct {
	AuthUsecase    auc.AuthUsecase
	ChoreUsecase   chuc.ChoreUsecase
	SessionUsecase suc.SessionUsecase
	PlayerUsecase  puc.PlayerUsecase
}

type Middlewares struct {
	ChatbotMiddleware chatbot.ChatbotMiddleware
}

type RestHandlers struct {
	AuthRestHandler ah.AuthRestHandler
}

type ChatbotHandlers struct {
	AuthChatbotHandler    ah.AuthChatbotHandler
	ChoreChatbotHandler   chh.ChoreChatbotHandler
	SessionChatbotHandler sh.SessionChatbotHandler
	PlayerChatbotHandler  ph.PlayerChatbotHandler
}

type RestService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	RestHandlers   *RestHandlers
	Usecases       *Usecases
	Middlewares    *Middlewares
}

type ChatbotListenerService struct {
	Middlewares    *Middlewares
	Handlers       *ChatbotHandlers
	Usecases       *Usecases
	Infrastructure *Infrastructure
	Config         *config.MainConfig
}

type Integration struct {
	SpotifyApiCall     spotify_api.SpotifyApiCallIntegration
	TelegramBotManager telegram_bot.TelegramBotIntegration
}
