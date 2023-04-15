package usecase

import (
	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

type Module struct {
	tokenRepository      tr.TokenRepository
	userRepository       ur.UserRepository
	spotifyPlayerApiCall spotify_api.SpotifyApiCallIntegration
	chatbotManager       telegram_bot.TelegramBotIntegration
}

type Opts struct {
	TokenRepository      tr.TokenRepository
	UserRepository       ur.UserRepository
	SpotifyPlayerApiCall spotify_api.SpotifyApiCallIntegration
	ChatbotManager       telegram_bot.TelegramBotIntegration
}

func New(o Opts) *Module {
	return &Module{
		tokenRepository:      o.TokenRepository,
		userRepository:       o.UserRepository,
		spotifyPlayerApiCall: o.SpotifyPlayerApiCall,
		chatbotManager:       o.ChatbotManager,
	}
}
