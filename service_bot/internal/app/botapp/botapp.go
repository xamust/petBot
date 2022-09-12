package botapp

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Понедельник", "Понедельник"),
		tgbotapi.NewInlineKeyboardButtonData("Вторник", "Вторник"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Среда", "Среда"),
		tgbotapi.NewInlineKeyboardButtonData("Четверг", "Четверг"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Пятница", "Пятница"),
		tgbotapi.NewInlineKeyboardButtonData("Суббота", "Суббота"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Воскресенье", "Воскресенье"),
	),
)

var clubKeyboard = make([][]tgbotapi.InlineKeyboardButton, 0)

// check callback and b.appendClubs
var clubsKeyboard = tgbotapi.NewInlineKeyboardMarkup(clubKeyboard...)

type BotApp struct {
	bot    *tgbotapi.BotAPI
	config *Config
	logger *logrus.Logger
}

func NewBot(config *Config) *BotApp {
	return &BotApp{
		config: config,
		logger: logrus.New(),
	}
}

func (b *BotApp) Start() error {

	//configure logger
	if err := b.configureLogger(); err != nil {
		return err
	}

	//configure bot...
	if err := b.configureBot(); err != nil {
		b.logger.Fatalln(err)
	}

	return nil
}

// config logger...
func (b *BotApp) configureLogger() error {
	//get level for logrus from configs...
	level, err := logrus.ParseLevel(b.config.LogLevel)
	if err != nil {
		return err
	}
	//set level for logrus...
	b.logger.SetLevel(level)
	return nil
}

// config bot...
func (b *BotApp) configureBot() error {
	bot, err := tgbotapi.NewBotAPI(b.config.APIKey)
	if err != nil {
		return err
	}
	b.bot = bot

	bot.Debug = b.config.BotDebug
	b.logger.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = b.config.BotTimeout
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "status":
				msg.Text = "I'm ok."
			case "fitness":
				msg.ReplyMarkup = numericKeyboard
				msg.Text = "Выбери день занятий:"
			}
			bot.Send(msg)
		} else if update.CallbackQuery != nil {

			//	//Respond to the callback query, telling Telegram to show the user
			//	// a message with the data received.
			//	callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			//	if _, err := bot.Request(callback); err != nil {
			//		panic(err)
			//	}

			//	// And finally, send a message containing the data received.
			//	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			//	sb := strings.Builder{}
			//	collector := fitness.CollectorInit()
			//	if _, err := collector.Search(); err != nil {
			//		b.logger.Error(err)
			//		msg.Text = err.Error()
			//	}
			//	for _, i2 := range collector.GetData(msg.Text) {
			//		sb.WriteString(i2)
			//	}
			//	msg.Text = sb.String()
			//	sb.Reset()

			//	v, err := bot.Send(msg)
			//	if err != nil {
			//		panic(err)
			//	}
			//	go b.delayDelete(time.Second*60, msg.ChatID, update.CallbackQuery.Message.MessageID)
			//	go b.delayDelete(time.Second*60, msg.ChatID, v.MessageID)
		}
	}
	return nil
}
