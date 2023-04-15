package handler

import (
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type ChoreChatbotHandler interface {
	HandleInitConversation(ci cbm.ChatInfo)
	HandleWelcome(ci cbm.ChatInfo)
}
