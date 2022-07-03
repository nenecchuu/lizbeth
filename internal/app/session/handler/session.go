package handler

import (
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type SessionChatbotHandler interface {
	HandleCreateSession(ci gm.ChatInfo)
}
