package handler

import (
	"github.com/gofiber/fiber/v2"
	gm "github.com/nenecchuu/lizbeth-be-core/internal/model"
)

type AuthRestHandler interface {
	HandleLinkageCallback(fc *fiber.Ctx) error
}

type AuthChatbotHandler interface {
	HandleHostAuthentication(ci gm.ChatInfo)
}
