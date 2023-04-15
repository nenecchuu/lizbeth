package handler

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/rs/zerolog/log"
)

func (x *ChatbotModule) HandleHostAuthentication(ci cbm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleHostAuthentication", nil)
	defer span.End()

	err := x.authUsecase.ProcessHostAuthentication(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
