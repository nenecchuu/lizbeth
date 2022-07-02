package model

import "github.com/nenecchuu/lizbeth-be-core/internal/constants"

type ChatInfo struct {
	Channel        constants.ChatbotChannelEnum
	SenderFullName string
	SenderId       string
	ChatId         string
	MessageId      string
	Message        string
}
