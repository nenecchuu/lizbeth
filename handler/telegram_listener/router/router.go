package router

import (
	"strings"

	"github.com/nenecchuu/lizbeth-be-core/init/service"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type Options struct {
	Handlers *service.ChatbotHandlers
}

type ChatbotRouter struct {
	handlers *service.ChatbotHandlers
}

func New(o *Options) *ChatbotRouter {
	return &ChatbotRouter{
		handlers: o.Handlers,
	}
}

func (c *ChatbotRouter) HandleMessage(message string, ci gm.ChatInfo) {
	var (
		ms     = strings.Split(message, ":")
		cmd    = ms[0]
		cmdval string
	)

	if len(ms) > 1 {
		cmdval = ms[1]
	}

	ci.Message = message

	switch cmd {
	case "auth":
		// c.handlers.AuthChatbotHandler.HandleLinkageCallback(ci)
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
	default:
		c.handlers.ChoreChatbotHandler.HandleWelcome(ci)
	}
}
