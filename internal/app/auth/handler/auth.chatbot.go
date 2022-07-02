package handler

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
	"github.com/rs/zerolog/log"
)

func (x *ChatbotModule) HandleHostAuthentication(ci gm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "auth.chatbot.HandleHostAuthentication", nil)
	defer span.End()

	err := x.authUsecase.ProcessHostAuthentication(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
