package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	auc "github.com/nenecchuu/lizbeth-be-core/internal/app/auth/usecase"
	cch "github.com/nenecchuu/lizbeth-be-core/internal/app/chore/usecase"
)

func (i *Initiator) InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories, integration *service.Integration) *service.Usecases {
	auth := auc.New(auc.Opts{
		TokenRepository:    repos.TokenRepository,
		UserRepository:     repos.UserRepository,
		SpotifyAuthApiCall: integration.SpotifyApiCall,
		ChatbotManager:     integration.TelegramBotManager,
	})

	chore := cch.New(cch.Opts{
		UserRepository: repos.UserRepository,
		ChatbotManager: integration.TelegramBotManager,
	})

	return &service.Usecases{
		AuthUsecase:  auth,
		ChoreUsecase: chore,
	}
}
