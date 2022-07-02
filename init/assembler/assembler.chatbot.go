package assembler

import (
	"github.com/nenecchuu/lizbeth-be-core/handler/telegram_listener"
	"github.com/nenecchuu/lizbeth-be-core/init/service"
)

func (a *assembler) assembleChatbotListener(service *service.ChatbotListenerService) telegram_listener.TelegramListenerHandler {
	server := a.NewChatbotListenerHandler(&telegram_listener.Opts{
		TelegramBotApi: service.Infrastructure.ChatbotModule,
		Handlers:       service.Handlers,
		Timeout:        service.Config.TelegramBotApi.Timeout,
		UpdateOffset:   service.Config.TelegramBotApi.UpdateOffset,
		DebugMode:      service.Config.TelegramBotApi.DebugMode,
	})

	return server
}
func (a *assembler) runTelegramListener() {
	a.chatbotListenerHanlder.Run()
}
