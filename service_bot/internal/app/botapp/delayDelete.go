package botapp

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

func (b *BotApp) delayDelete(d time.Duration, chatID int64, messageID int) {
	time.Sleep(d)
	delMess := tgbotapi.NewDeleteMessage(chatID, messageID)
	b.bot.Send(delMess)
}
