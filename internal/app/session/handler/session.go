package handler

import (
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type SessionChatbotHandler interface {
	HandleCreateSession(ci cbm.ChatInfo)
}
