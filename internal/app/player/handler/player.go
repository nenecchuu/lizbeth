package handler

import (
	gam "github.com/nenecchuu/lizbeth-be-core/internal/model/auth"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type PlayerChatbotHandler interface {
	HandleQueueTrack(ci cbm.ChatInfo, cm *gam.CommandMetadata)
}
