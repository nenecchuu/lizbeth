package chatbot

import (
	"context"

	gam "github.com/nenecchuu/lizbeth-be-core/internal/model/auth"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type ChatbotMiddleware interface {
	ParseAndValidateSenderData(ctx context.Context, ci cbm.ChatInfo) (res *gam.CommandMetadata, err error)
}
