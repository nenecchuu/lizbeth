package usecase

import (
	"github.com/nenecchuu/arcana/util"
	sr "github.com/nenecchuu/lizbeth-be-core/internal/app/session/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type Module struct {
	sessionRepository  sr.SessionRepository
	userRepository     ur.UserRepository
	spotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	chatbotManager     telegram_bot.TelegramBotIntegration
	snowflakeManager   *util.POD
}

type Opts struct {
	UserRepository     ur.UserRepository
	SessionRepository  sr.SessionRepository
	SpotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	ChatbotManager     telegram_bot.TelegramBotIntegration
	SnowflakeManager   *util.POD
}

func New(o Opts) *Module {
	return &Module{
		sessionRepository:  o.SessionRepository,
		userRepository:     o.UserRepository,
		spotifyAuthApiCall: o.SpotifyAuthApiCall,
		chatbotManager:     o.ChatbotManager,
		snowflakeManager:   o.SnowflakeManager,
	}
}
