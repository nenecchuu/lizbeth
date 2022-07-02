package repository

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/app/token/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (x *Module) StoreToken(ctx context.Context, token *model.TokenNoSqlSchema) error {
	ctx, span := tracer.StartSpan(ctx, "api_call.spotify.GetUserInfo", nil)
	defer span.End()

	var (
		e error
		c *mongo.Collection
	)

	c = x.mongoManager.Collection(model.TokenCollectionName)

	_, e = c.InsertOne(ctx, token)

	if e != nil {
		return e
	}

	return nil
}
