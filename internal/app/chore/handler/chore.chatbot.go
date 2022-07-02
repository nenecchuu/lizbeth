package handler

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
	"github.com/rs/zerolog/log"
)

func (x *ChatbotModule) HandleInitConversation(ci gm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleLinkageCallback", nil)
	defer span.End()

	err := x.choreUsecase.ProcessInitConversation(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}
}

func (x *ChatbotModule) HandleWelcome(ci gm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleLinkageCallback", nil)
	defer span.End()

	err := x.choreUsecase.ProcessWelcome(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}
}
