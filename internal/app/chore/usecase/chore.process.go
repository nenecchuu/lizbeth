package usecase

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/rs/zerolog/log"
)

func (x *Module) ProcessInitConversation(ctx context.Context, ci cbm.ChatInfo) error {
	ctx, span := tracer.StartSpan(ctx, "chore.uc.ProcessInitConversation", nil)
	defer span.End()

	var (
		err   error
		udata *um.UserNoSqlSchema
	)

	udata, err = x.userRepository.FindUserByChatbotUserId(ctx, ci.SenderId, ci.Channel)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	if udata == nil {
		udata = udata.BuildFromChatInfo(ci)

		err = x.userRepository.StoreUser(ctx, udata)
		if err != nil {
			log.Err(err).Msg(err.Error())
			return err
		}
	}

	err = x.chatbotManager.SendProcessInitConversationMessage(ctx, ci.ChatId)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (x *Module) ProcessWelcome(ctx context.Context, ci cbm.ChatInfo) error {
	ctx, span := tracer.StartSpan(ctx, "chore.uc.ProcessWelcome", nil)
	defer span.End()

	err := x.chatbotManager.SendWelcomeConversationMessage(ctx, ci.ChatId)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return nil
}
