package usecase

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

func (x *Module) ProcessQueueTrack(ctx context.Context, ci cbm.ChatInfo) (err error) {
	ctx, span := tracer.StartSpan(ctx, "player.uc.ProcessQueueTrack", nil)
	defer span.End()

	var (
		keyword = ci.Message
	)

	x.spotifyPlayerApiCall.SearchTracks(ctx, "", keyword)

	return err
}
