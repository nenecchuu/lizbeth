package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	auc "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/usecase"
	cuc "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/usecase"
	suc "github.com/nenecchuu/lizbeth-be-core/internal/app/session/usecase"
)

func (i *Initiator) InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories, integration *service.Integration) *service.Usecases {
	auth := auc.New(auc.Opts{
		TokenRepository:    repos.TokenRepository,
		UserRepository:     repos.UserRepository,
		SpotifyAuthApiCall: integration.SpotifyApiCall,
		ChatbotManager:     integration.TelegramBotManager,
	})

	chore := cuc.New(cuc.Opts{
		UserRepository: repos.UserRepository,
		ChatbotManager: integration.TelegramBotManager,
	})

	session := suc.New(suc.Opts{
		UserRepository:     repos.UserRepository,
		SessionRepository:  repos.SessionRepository,
		ChatbotManager:     integration.TelegramBotManager,
		SpotifyAuthApiCall: integration.SpotifyApiCall,
		SnowflakeManager:   infra.Snowflake,
	})

	return &service.Usecases{
		AuthUsecase:    auth,
		ChoreUsecase:   chore,
		SessionUsecase: session,
	}
}
