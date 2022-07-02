package constants

import (
	"fmt"

	gc "github.com/nenecchuu/lizbeth-be-core/internal/constants"
	"github.com/nenecchuu/lizbeth-be-core/internal/integration/telegram_bot/model"
)

var (
	WelcomeMessageRows = []model.TelegramInlineKeyboardRow{
		{
			Data: []model.TelegramInlineKeyboardData{
				{
					Value: string(gc.ChatbotCommandMessageStart),
					Text:  "Start",
				},
			},
		},
	}
	InitConversationRows = []model.TelegramInlineKeyboardRow{
		{
			Data: []model.TelegramInlineKeyboardData{
				{
					Value: fmt.Sprintf("%s:%s", gc.ChatbotCommandMessageEnterRole, gc.ChatbotUserRoleHost),
					Text:  "Host",
				},
			},
		}, {
			Data: []model.TelegramInlineKeyboardData{
				{
					Value: fmt.Sprintf("%s:%s", gc.ChatbotCommandMessageEnterRole, gc.ChatbotUserRoleHost),
					Text:  "Guest",
				},
			},
		},
	}
	HostActionsRows = []model.TelegramInlineKeyboardRow{
		{
			Data: []model.TelegramInlineKeyboardData{
				{
					Value: fmt.Sprintf("%s:%s", gc.ChatbotCommandMessageSession, gc.ChatbotSessionCommandInit),
					Text:  "Create New Session",
				},
			},
		},
	}
)
