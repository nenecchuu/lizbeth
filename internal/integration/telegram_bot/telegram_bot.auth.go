package telegram_bot

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot/constants"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot/model"
	"github.com/rs/zerolog/log"
)

func (x *Module) SendInitializeLinkageMessage(ctx context.Context, chat_id string, linkage_url string) error {
	ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendWelcomeMessage", nil)
	defer span.End()

	initLinkageRows := []model.TelegramInlineKeyboardRow{
		{
			Data: []model.TelegramInlineKeyboardData{
				{
					URL:  linkage_url,
					Text: "Login to Spotify",
				},
			},
		},
	}

	err := x.sendInlineTextMessage(ctx, constants.InitLinkageMessage, initLinkageRows, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}

func (x *Module) SendLinkageSuccessMessage(ctx context.Context, chat_id string) error {
	ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendWelcomeMessage", nil)
	defer span.End()

	err := x.sendTextMessage(ctx, "Linkage Success", chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}
