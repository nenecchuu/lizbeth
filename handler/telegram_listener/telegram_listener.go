package telegram_listener

import (
	"context"
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/nenecchuu/lizbeth-be-core/handler/telegram_listener/router"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
	"github.com/nenecchuu/lizbeth-be-core/internal/constants"
	cbm "github.com/nenecchuu/lizbeth-be-core/internal/model/chatbot"
	"github.com/rs/zerolog/log"
)

type Opts struct {
	TelegramBotApi *tgbotapi.BotAPI
	Handlers       *service.ChatbotHandlers
	DebugMode      bool
	Timeout        int
	UpdateOffset   int
}

type Handler struct {
	telegramBotApi *tgbotapi.BotAPI
	router         *router.ChatbotRouter
	debugMode      bool
	timeout        int
	updateOffset   int
}

type TelegramListenerHandler interface {
	Run()
}

func New(o *Opts) TelegramListenerHandler {
	router := router.New(&router.Options{
		Handlers: o.Handlers,
	})

	return &Handler{
		telegramBotApi: o.TelegramBotApi,
		router:         router,
		debugMode:      o.DebugMode,
		timeout:        o.Timeout,
		updateOffset:   o.UpdateOffset,
	}
}

func (h *Handler) Run() {
	ctx := context.Background()
	if h.debugMode {
		h.telegramBotApi.Debug = true

		log.Printf("Authorized on account %s", h.telegramBotApi.Self.UserName)
	}

	u := tgbotapi.NewUpdate(h.updateOffset)
	u.Timeout = h.timeout

	updates := h.telegramBotApi.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			ci := cbm.ChatInfo{
				Channel:        constants.ChatbotChannelTelegram,
				SenderFullName: fmt.Sprintf("%s %s", update.Message.Chat.FirstName, update.Message.Chat.LastName),
				SenderId:       update.Message.Chat.UserName,
				ChatId:         strconv.FormatInt(update.Message.Chat.ID, 10),
				MessageId:      strconv.Itoa(update.Message.MessageID),
			}
			h.router.HandleMessage(ctx, update.Message.Text, ci)

		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			_, err := h.telegramBotApi.Request(callback)

			if err != nil {
				log.Err(err).Msg(err.Error())
			}

			if err == nil {
				ci := cbm.ChatInfo{
					Channel:        constants.ChatbotChannelTelegram,
					SenderFullName: fmt.Sprintf("%s %s", update.CallbackQuery.Message.Chat.FirstName, update.CallbackQuery.Message.Chat.LastName),
					SenderId:       update.CallbackQuery.Message.Chat.UserName,
					ChatId:         strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10),
					MessageId:      strconv.Itoa(update.CallbackQuery.Message.MessageID),
				}

				h.router.HandleMessage(ctx, update.CallbackQuery.Data, ci)
			}
		}
	}
}
