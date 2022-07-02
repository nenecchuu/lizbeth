package usecase

import (
	"context"

	"github.com/nenecchuu/arcana/tracer"
	"github.com/nenecchuu/lizbeth-be-core/internal/app/auth/model"
)

// TODO: implement transaction
func (x *Module) ProcessCreateNewSession(ctx context.Context, data *model.LinkageCallback) error {
	ctx, span := tracer.StartSpan(ctx, "session.uc.ProcessCreateNewSession", nil)
	defer span.End()

	return nil
}
