package handler

import (
	"github.com/gofiber/fiber/v2"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
)

type AuthRestHandler interface {
	HandleLinkageCallback(fc *fiber.Ctx) error
}

type AuthChatbotHandler interface {
	HandleHostAuthentication(ci cbm.ChatInfo)
}
