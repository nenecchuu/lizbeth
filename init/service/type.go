package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nenecchuu/arcana/mongo"
	"github.com/nenecchuu/arcana/util"
	"github.com/nenecchuu/lizbeth-be-core/config"

	ah "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/handler"
	auc "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/usecase"

	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"

	chh "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/handler"
	chuc "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/usecase"

	sh "github.com/nenecchuu/lizbeth-be-core/internal/app/session/handler"
	sr "github.com/nenecchuu/lizbeth-be-core/internal/app/session/repository"
	suc "github.com/nenecchuu/lizbeth-be-core/internal/app/session/usecase"

	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type Infrastructure struct {
	ChatbotModule *tgbotapi.BotAPI
	Mongo         *mongo.Database
	Snowflake     *util.POD
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
}

type Middlewares struct {
}

type RestHandlers struct {
	AuthRestHandler ah.AuthRestHandler
}

type ChatbotHandlers struct {
	AuthChatbotHandler    ah.AuthChatbotHandler
	ChoreChatbotHandler   chh.ChoreChatbotHandler
	SessionChatbotHandler sh.SessionChatbotHandler
}

type RestService struct {
	Config         *config.MainConfig
	Infrastructure *Infrastructure
	RestHandlers   *RestHandlers
	Usecases       *Usecases
	Middlewares    *Middlewares
}

type ChatbotListenerService struct {
	Handlers       *ChatbotHandlers
	Usecases       *Usecases
	Infrastructure *Infrastructure
	Config         *config.MainConfig
}

type Integration struct {
	SpotifyApiCall     spotify_api.SpotifyApiCallIntegration
	TelegramBotManager telegram_bot.TelegramBotIntegration
}
