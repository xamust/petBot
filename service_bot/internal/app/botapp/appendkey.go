package botapp

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *BotApp) appendClubs(text map[string]string) tgbotapi.InlineKeyboardMarkup {
	count := 0
	keyButt := make([]tgbotapi.InlineKeyboardButton, 0)
	for i, i2 := range text {
		if count == 2 {
			clubKeyboard = append(clubKeyboard, tgbotapi.NewInlineKeyboardRow(keyButt...))
			keyButt = make([]tgbotapi.InlineKeyboardButton, 0)
			count = 0
		}
		keyButt = append(keyButt, tgbotapi.NewInlineKeyboardButtonData(i, i2))
		count++
	}
	return tgbotapi.NewInlineKeyboardMarkup(clubKeyboard...)
}

func (b *BotApp) appendInClubs() {

}
