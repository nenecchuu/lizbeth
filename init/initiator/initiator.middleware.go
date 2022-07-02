package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
)

func (i *Initiator) InitMiddleware(cfg *config.MainConfig, infra *service.Infrastructure, repos *service.Repositories) *service.Middlewares {

	return &service.Middlewares{}
}
