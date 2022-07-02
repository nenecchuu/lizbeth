package handler

import (
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type ChoreChatbotHandler interface {
	HandleInitConversation(ci gm.ChatInfo)
	HandleWelcome(ci gm.ChatInfo)
}
