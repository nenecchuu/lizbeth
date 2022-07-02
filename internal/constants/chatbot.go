package constants

type ChatbotChannelEnum int

const (
	ChatbotChannelTelegram ChatbotChannelEnum = 1
)

type ChatbotCommandMessage string

const (
	ChatbotCommandMessageStart     ChatbotCommandMessage = "start"
	ChatbotCommandMessageEnterRole ChatbotCommandMessage = "enter_role_as"
	ChatbotCommandMessageSession   ChatbotCommandMessage = "session"
)

type ChatbotUserRole string

const (
	ChatbotUserRoleGuest ChatbotUserRole = "guest"
	ChatbotUserRoleHost  ChatbotUserRole = "host"
)

type ChatbotSessionCommand string

const (
	ChatbotSessionCommandInit ChatbotSessionCommand = "init"
)
