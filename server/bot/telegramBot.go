package bot

import (
	"database/sql"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"makebex-backend/server/bot/config"
	"makebex-backend/server/bot/model"
	"makebex-backend/server/bot/views/content/en"
)

var bot *tgbotapi.BotAPI

func initBot() {
	var err error

	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: use when dev.
	bot.Debug = true
}

func StartBot(db *sql.DB) {
	initBot()

	bot.RemoveWebhook()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		Handle(update, db)
	}
}

func StartBotWithWebhook(db *sql.DB) {
	var err error

	initBot()
	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://cb76a226.ngrok.io/"+bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	//go http.ListenAndServe("0.0.0.0:10000",  nil)

	go func() {
		for update := range updates {
			if update.Message == nil {
				continue
			}

			Handle(update, db)
		}
	}()

}

func Handle(update tgbotapi.Update, db *sql.DB) {
	if update.Message.Contact != nil {
		var msg tgbotapi.MessageConfig

		user, err := model.CreateUser(db, update)

		if err != nil {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, en.Error)
		} else {
			text := fmt.Sprintf(en.Welcome, user.ID, user.Username)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ParseMode = tgbotapi.ModeMarkdown
		}

		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, en.SendPhoneMessageText)
	msg.ParseMode = tgbotapi.ModeMarkdown
	//msg.ReplyToMessageID = update.Message.MessageID

	var contact = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact(en.SendPhone),
		),
	)

	msg.ReplyMarkup = contact

	bot.Send(msg)
}