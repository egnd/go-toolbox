// Package tgchain has chained wrapper for handling tg updates
package tgchain

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// ITgAPI is interface for Telegram API instance.
type ITgAPI interface {
	GetUpdatesChan(tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	Send(tgbotapi.Chattable) (tgbotapi.Message, error)
	GetMe() (tgbotapi.User, error)
	GetChatMember(tgbotapi.ChatConfigWithUser) (tgbotapi.ChatMember, error)
	LeaveChat(tgbotapi.ChatConfig) (tgbotapi.APIResponse, error)
	AnswerInlineQuery(tgbotapi.InlineConfig) (tgbotapi.APIResponse, error)
}

// WarnUpd is a method for logging warnings.
type WarnUpd func(string, tgbotapi.Update)

// ErrUpd is a method for logging errors.
type ErrUpd func(string, tgbotapi.Update, error)
