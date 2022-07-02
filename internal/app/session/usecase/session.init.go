package usecase

import (
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/session/repository"
	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type Module struct {
	tokenRepository    tr.TokenRepository
	sessionRepository  ur.SessionRepository
	spotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	chatbotManager     telegram_bot.TelegramBotIntegration
}

type Opts struct {
	TokenRepository    tr.TokenRepository
	SessionRepository  ur.SessionRepository
	SpotifyAuthApiCall spotify_api.SpotifyApiCallIntegration
	ChatbotManager     telegram_bot.TelegramBotIntegration
}

func New(o Opts) *Module {
	return &Module{
		tokenRepository:    o.TokenRepository,
		sessionRepository:  o.SessionRepository,
		spotifyAuthApiCall: o.SpotifyAuthApiCall,
		chatbotManager:     o.ChatbotManager,
	}
}
