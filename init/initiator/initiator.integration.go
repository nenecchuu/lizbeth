package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/spotify_api"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot"
)

func (i *Initiator) InitIntegration(cfg *config.MainConfig, infra *service.Infrastructure) *service.Integration {
	sa := spotify_api.New(spotify_api.Opts{
		SpotifyConfig: &cfg.Spotify,
	})

	tgManager := telegram_bot.New(telegram_bot.Opts{
		TelegramBotApi: infra.ChatbotModule,
	})

	return &service.Integration{
		SpotifyApiCall:     sa,
		TelegramBotManager: tgManager,
	}
}
