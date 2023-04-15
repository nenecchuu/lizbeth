package router

import (
	"context"
	"strings"

	"github.com/nenecchuu/lizbeth-be-core/init/service"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type Options struct {
	Middleware *service.Middlewares
	Handlers   *service.ChatbotHandlers
}

type ChatbotRouter struct {
	middlewares *service.Middlewares
	handlers    *service.ChatbotHandlers
}

func New(o *Options) *ChatbotRouter {
	return &ChatbotRouter{
		handlers: o.Handlers,
	}
}

func (c *ChatbotRouter) HandleMessage(ctx context.Context, message string, ci cbm.ChatInfo) {
	var (
		ms     = strings.Split(message, " ")
		cmd    = ms[0]
		cmdval string
		// err    error
	)

	if len(ms) > 1 {
		cmdval = ms[1]
	}

	ci.Message = message

	switch cmd {
	case string(constants.ChatbotCommandMessageStart):
		c.handlers.ChoreChatbotHandler.HandleInitConversation(ci)
	case string(constants.ChatbotCommandMessageEnterRole):
		switch cmdval {
		case string(constants.ChatbotUserRoleHost):
			c.handlers.AuthChatbotHandler.HandleHostAuthentication(ci)
		case string(constants.ChatbotUserRoleGuest):
			c.handlers.ChoreChatbotHandler.HandleInitConversation(ci)
		}
	case string(constants.ChatbotCommandMessageCreateSession):
		c.handlers.SessionChatbotHandler.HandleCreateSession(ci)
	case string(constants.ChatbotCommandMessageQueueTrack):
		cm, err := c.middlewares.ChatbotMiddleware.ParseAndValidateSenderData(ctx, ci)
		if err == nil {
			c.handlers.PlayerChatbotHandler.HandleQueueTrack(ci, cm)
		}
	default:
		c.handlers.ChoreChatbotHandler.HandleWelcome(ci)
	}
}
