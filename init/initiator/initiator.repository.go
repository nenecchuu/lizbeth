package initiator

import (
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	sr "github.com/nenecchuu/lizbeth-be-core/internal/app/session/repository"
	tr "github.com/nenecchuu/lizbeth-be-core/internal/app/token/repository"
	ur "github.com/nenecchuu/lizbeth-be-core/internal/app/user/repository"
)

func (i *Initiator) InitRepository(cfg *config.MainConfig, infra *service.Infrastructure) *service.Repositories {
	token := tr.New(tr.Opts{
		MongoManager: infra.Mongo.Database,
	})
	user := ur.New(ur.Opts{
		MongoManager: infra.Mongo.Database,
	})

	session := sr.New(sr.Opts{
		MongoManager: infra.Mongo.Database,
	})

	return &service.Repositories{
		TokenRepository:   token,
		UserRepository:    user,
		SessionRepository: session,
	}
}
