package handler

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	gam "github.com/nenecchuu/lizbeth-be-core/internal/model/auth"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/rs/zerolog/log"
)

func (x *ChatbotModule) HandleQueueTrack(ci cbm.ChatInfo, cm *gam.CommandMetadata) {
	ctx, span := tracer.StartSpan(context.Background(), "player.chatbot.HandleQueueTrack", nil)
	defer span.End()

	err := x.playerUsecase.ProcessQueueTrack(ctx, ci)

	if err != nil {
		log.Err(err).Msg(err.Error())
	}

}
