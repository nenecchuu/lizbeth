package repository

import (
	"context"
	"time"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/app/session/model"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (x *Module) StoreSession(ctx context.Context, session *model.SessionNoSqlSchema) error {
	ctx, span := tracer.StartSpan(ctx, "repo.session.GetSessionInfo", nil)
	defer span.End()

	if session.Id == primitive.NilObjectID {
		session.Id = primitive.NewObjectIDFromTimestamp(time.Now())
	}

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	_, e = c.InsertOne(ctx, session)

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) UpdateSession(ctx context.Context, id primitive.ObjectID, session *model.SessionNoSqlSchema) error {
	ctx, span := tracer.StartSpan(ctx, "repo.session.GetSessionInfo", nil)
	defer span.End()

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	_, e = c.UpdateByID(ctx, id, bson.M{"$set": session})

	if e != nil {
		return e
	}

	return nil
}

func (x *Module) FindSessionById(ctx context.Context, id primitive.ObjectID) (*model.SessionNoSqlSchema, error) {
	ctx, span := tracer.StartSpan(ctx, "repo.session.FindSessionById", nil)
	defer span.End()

	var (
		e    error
		c    *mongo.Collection
		f    = bson.D{{"_id", id}}
		res  *mongo.SingleResult
		data = &model.SessionNoSqlSchema{}
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	res = c.FindOne(ctx, f)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, res.Err()
	}

	e = res.Decode(data)
	if e != nil {
		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return data, nil
}

func (x *Module) FindSessionByChatbotSessionId(ctx context.Context, chatbotSessionId string, chatbotChannel constants.ChatbotChannelEnum) (*model.SessionNoSqlSchema, error) {
	ctx, span := tracer.StartSpan(ctx, "repo.session.FindSessionByChatbotSessionId", nil)
	defer span.End()

	var (
		e    error
		c    *mongo.Collection
		f    = bson.D{{"chatbot_session_id", chatbotSessionId}, {"chatbot_channel", chatbotChannel}}
		res  *mongo.SingleResult
		data = &model.SessionNoSqlSchema{}
	)

	c = x.mongoManager.Collection(model.SessionCollectionName)

	res = c.FindOne(ctx, f)

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, res.Err()
	}

	e = res.Decode(data)
	if e != nil {
		log.Err(e).Msg(e.Error())
		return nil, e
	}

	return data, nil
}
