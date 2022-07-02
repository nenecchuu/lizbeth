package handler

import (
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type AuthChatbotHandler interface {
	HandleHostAuthentication(ci gm.ChatInfo)
}
