package botapp

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *BotApp) appendClubs(text, data string) {
	clubKeyboard = append(clubKeyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(text, data),
	))
}
