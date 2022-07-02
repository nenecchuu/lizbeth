package telegram_listener

import (
	"log"

	"github.com/nenecchuu/arcana/env"
	"github.com/nenecchuu/arcana/logger"
	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/init/assembler"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	serveTelegramListenerCmd = &cobra.Command{
		Use:              "serve-telegram-listener",
		Short:            "Lizbeth Telegram Listener",
		PersistentPreRun: rootPreRun,
		RunE:             runTelegramListener,
	}
)

func rootPreRun(cmd *cobra.Command, args []string) {
	logger.InitGlobalLogger(&logger.Config{
		ServiceName: "lizbeth-telegram-listener",
		Level:       zerolog.DebugLevel,
	})
}

func ServeTelegramCmd() *cobra.Command {
	return serveTelegramListenerCmd
}

func runTelegramListener(cmd *cobra.Command, args []string) error {
	configURL, _ := cmd.Flags().GetString("config")
	bootstrapTelegramListener(assembler.NewAssembler(), configURL)
	return nil
}

func bootstrapTelegramListener(starter assembler.AssemblerManager, configPath string) {
	err := tracer.Init(&tracer.TracerConfig{
		UseJaeger:   false,
		Environment: env.GetEnvironmentName(),
		ServiceName: "lizbeth-telegram-listener",
	})

	if err != nil {
		log.Fatalln(err)
	}

	starter = starter.BuildService(configPath).AssembleTelegramListernerApplication()
	starter.RunTelegramListenerApplication()
}
