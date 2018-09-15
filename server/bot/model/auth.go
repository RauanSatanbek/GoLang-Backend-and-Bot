package model

import (
	"database/sql"
	"gopkg.in/telegram-bot-api.v4"
	"makebex-backend/server/auth"
)

func CreateUser(db *sql.DB, update tgbotapi.Update) (auth.User, error ) {
	user := auth.User{
		Username: update.Message.From.UserName,
		TelegramID: update.Message.Contact.UserID,
		FirstName: update.Message.Contact.FirstName,
		PhoneNumber: update.Message.Contact.PhoneNumber,
	}

	err := user.Create(db)


	return user, err
}