package telegram_bot

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot/constants"
	"github.com/rs/zerolog/log"
)

func (x *Module) SendProcessInitConversationMessage(ctx context.Context, chat_id string) error {
	ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendProcessInitConversationMessage", nil)
	defer span.End()

	err := x.sendInlineTextMessage(ctx, constants.InitConversationMessage, constants.InitConversationRows, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}

func (x *Module) SendWelcomeConversationMessage(ctx context.Context, chat_id string) error {
	ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendWelcomeMessage", nil)
	defer span.End()

	err := x.sendInlineTextMessage(ctx, constants.WelcomeMessage, constants.WelcomeMessageRows, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}

func (x *Module) SendHostActionsMessage(ctx context.Context, chat_id string) error {
	ctx, span := tracer.StartSpan(ctx, "chatbot.telegram.SendWelcomeMessage", nil)
	defer span.End()

	err := x.sendTextMessage(ctx, constants.HostActionsMessage, chat_id)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

	return nil
}
