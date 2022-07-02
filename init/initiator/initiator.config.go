package initiator

import (
	"github.com/google/gops/agent"
	configpkg "github.com/nenecchuu/arcana/config"
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/rs/zerolog/log"
)

var (
	errInitConfig = "failed to initiate config"
)

//Main Config
func (i *Initiator) InitConfig(configPath string) *config.MainConfig {
	if err := i.agentListen(agent.Options{
		ShutdownCleanup: true, // automatically closes on os.Interrupt
	}); err != nil {
		log.Fatal().Err(err).Msg(errInitConfig)
	}

	cfg := &config.MainConfig{}
	log.Info().Msgf("reading config from %s", configPath)
	err := configpkg.ReadConfig(cfg, configPath, "config")
	if err != nil {
		log.Fatal().Err(err).Msg(errInitConfig)
	}

	return cfg

}
