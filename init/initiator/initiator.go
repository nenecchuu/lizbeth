package initiator

import (
	"flag"

	configpkg "github.com/nenecchuu/arcana/config"
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"

	"github.com/google/gops/agent"
)

type Init func()
type agentListen func(opts agent.Options) error
type ReadConfig func(cfg interface{}, path string, module string) error

type InitiatorManager interface {
	InitConfig(configPath string) *config.MainConfig
	InitInfrastructure(cfg *config.MainConfig) *service.Infrastructure
	InitRepository(cfg *config.MainConfig, infra *service.Infrastructure) *service.Repositories
	InitIntegration(cfg *config.MainConfig, infra *service.Infrastructure) *service.Integration
	InitUsecase(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories, integration *service.Integration) *service.Usecases
	InitMiddleware(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories) *service.Middlewares
	InitRestHandler(cfg *config.MainConfig, infra *service.Infrastructure, uc *service.Usecases) *service.RestHandlers
	InitChatbotHandler(cfg *config.MainConfig, infra *service.Infrastructure, integration *service.Integration, uc *service.Usecases) *service.ChatbotHandlers
	InitRestService(cfg *config.MainConfig, infra *service.Infrastructure, hdl *service.RestHandlers, uc *service.Usecases, mw *service.Middlewares) *service.RestService
	InitChatbotListenerService(cfg *config.MainConfig, hdl *service.ChatbotHandlers, infra *service.Infrastructure, uc *service.Usecases) *service.ChatbotListenerService
}

type Initiator struct {
	FlagParse   Init
	agentListen agentListen
	ReadConfig  ReadConfig
}

func New() InitiatorManager {
	return &Initiator{
		FlagParse:   flag.Parse,
		agentListen: agent.Listen,
		ReadConfig:  configpkg.ReadConfig,
	}
}

func (i *Initiator) InitRestService(cfg *config.MainConfig, infra *service.Infrastructure, hdl *service.RestHandlers, uc *service.Usecases, mw *service.Middlewares) *service.RestService {
	svc := service.RestService{
		Config:         cfg,
		Infrastructure: infra,
		RestHandlers:   hdl,
		Usecases:       uc,
		Middlewares:    mw,
	}

	return &svc
}

func (i *Initiator) InitChatbotListenerService(cfg *config.MainConfig, hdl *service.ChatbotHandlers, infra *service.Infrastructure, uc *service.Usecases) *service.ChatbotListenerService {
	svc := service.ChatbotListenerService{
		Config:         cfg,
		Infrastructure: infra,
		Usecases:       uc,
		Handlers:       hdl,
	}

	return &svc
}
