package handler

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
	"github.com/rs/zerolog/log"
)

func (x *ChatbotModule) HandleCreateSession(ci gm.ChatInfo) {
	ctx, span := tracer.StartSpan(context.Background(), "session.chatbot.HandleCreateSession", nil)
	defer span.End()

	err := x.sessionUsecase.ProcessCreateNewSession(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
