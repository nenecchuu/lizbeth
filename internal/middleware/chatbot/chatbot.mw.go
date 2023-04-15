package chatbot

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	sm "github.com/nenecchuu/lizbeth-be-core/internal/app/session/model"
	tm "github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
	gam "github.com/nenecchuu/lizbeth-be-core/internal/model/auth"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (x *ChatbotMiddlewareModule) ParseAndValidateSenderData(ctx context.Context, ci cbm.ChatInfo) (res *gam.CommandMetadata, err error) {
	ctx, span := tracer.StartSpan(ctx, "chatbot.mw.ParseAndValidateSenderData", nil)
	defer span.End()

	var (
		u *um.UserNoSqlSchema
		s *sm.SessionNoSqlSchema
		t *tm.TokenNoSqlSchema
	)

	u, err = x.userRepository.FindUserByChatbotUserId(ctx, ci.SenderId, ci.Channel)

	if err != nil {
		log.Err(err).Msgf(err.Error())
		x.chatbotManager.SendErrorMessage(ctx, ci.ChatId, err.Error())
		return nil, err
	}

	if u.ActiveSessionId != primitive.NilObjectID {
		s, err = x.sessionRepository.FindSessionById(ctx, u.ActiveSessionId)
		if err != nil {
			log.Err(err).Msgf(err.Error())
			x.chatbotManager.SendErrorMessage(ctx, ci.ChatId, err.Error())
			return nil, err
		}

		t, err = x.tokenRepository.FindAndValidateTokenByUserId(ctx, s.HostId)

		if err != nil {
			log.Err(err).Msgf(err.Error())
			x.chatbotManager.SendErrorMessage(ctx, ci.ChatId, err.Error())
			return nil, err
		}
	}

	res = &gam.CommandMetadata{
		User:      *u,
		Session:   *s,
		HostToken: *t,
	}

	return res, nil
}
