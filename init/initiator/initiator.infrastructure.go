package initiator

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nenecchuu/arcana/mongo"
	"github.com/nenecchuu/lizbeth-be-core/config"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	"github.com/rs/zerolog/log"
)

func (i *Initiator) InitInfrastructure(cfg *config.MainConfig) *service.Infrastructure {
	mongo := mongo.MongoConnectClient(&mongo.Client{
		URI:            cfg.Mongo.URI,
		DB:             cfg.Mongo.DB,
		AppName:        cfg.Mongo.DB,
		ConnectTimeout: time.Duration(cfg.Mongo.ConnectionTimeout) * time.Second,
		PingTimeout:    time.Duration(cfg.Mongo.PingTimeout) * time.Second,
	})

	chatbot := initChatbotApi(cfg)

	return &service.Infrastructure{
		Mongo:         mongo,
		ChatbotModule: chatbot,
	}
}

func initChatbotApi(cfg *config.MainConfig) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotApi.Token)

	if err != nil {
		log.Panic().Msg(err.Error())
	}

	return bot
}
