package handler

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/rs/zerolog/log"
)

func (x *ChatbotModule) HandleInitConversation(ci cbm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleLinkageCallback", nil)
	defer span.End()

	err := x.choreUsecase.ProcessInitConversation(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}
}

func (x *ChatbotModule) HandleWelcome(ci cbm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleLinkageCallback", nil)
	defer span.End()

	err := x.choreUsecase.ProcessWelcome(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}
}
