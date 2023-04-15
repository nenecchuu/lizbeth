package usecase

import (
	"context"
	"time"

	"github.com/nenecchuu/arcana/tracer"
	sm "github.com/nenecchuu/lizbeth-be-core/internal/app/session/model"
	um "github.com/nenecchuu/lizbeth-be-core/internal/app/user/model"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/nenecchuu/lizbeth-be-core/internal/util"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (x *Module) ProcessCreateNewSession(ctx context.Context, ci cbm.ChatInfo) error {
	ctx, span := tracer.StartSpan(ctx, "session.uc.ProcessCreateNewSession", nil)
	defer span.End()

	var (
		user *um.UserNoSqlSchema
		data *sm.SessionNoSqlSchema
		err  error
		now  = time.Now()
	)

	user, err = x.userRepository.FindUserByChatbotUserId(ctx, ci.SenderId, ci.Channel)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	if user == nil {
		return util.NewDataNotFoundErr("User")
	}

	data = &sm.SessionNoSqlSchema{
		Id:        primitive.NewObjectIDFromTimestamp(now),
		Code:      user.ChatbotUserId,
		HostId:    user.Id,
		CreatedAt: now,
		ExpireAt:  now.Add(time.Hour * 24),
	}

	err = x.sessionRepository.StoreSession(ctx, data)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	user.ActiveSessionId = data.Id
	err = x.userRepository.UpdateUser(ctx, user.Id, user)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	err = x.chatbotManager.SendSessionCreatedMessage(ctx, ci.ChatId, data.Code)

	if err != nil {
		log.Err(err).Msg(err.Error())
		return err
	}

	return err
}
