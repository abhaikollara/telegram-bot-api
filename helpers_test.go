package telegram_test

import (
	"testing"

	telegram "github.com/abhaikollara/telegram-bot-api"
)

func TestNewInlineQueryResultArticle(t *testing.T) {
	result := telegram.NewInlineQueryResultArticle("id", "title", "message")

	if result.Type != "article" ||
		result.ID != "id" ||
		result.Title != "title" ||
		result.InputMessageContent.(telegram.InputTextMessageContent).Text != "message" {
		t.Fail()
	}
}

func TestNewInlineQueryResultArticleMarkdown(t *testing.T) {
	result := telegram.NewInlineQueryResultArticleMarkdown("id", "title", "*message*")

	if result.Type != "article" ||
		result.ID != "id" ||
		result.Title != "title" ||
		result.InputMessageContent.(telegram.InputTextMessageContent).Text != "*message*" ||
		result.InputMessageContent.(telegram.InputTextMessageContent).ParseMode != "Markdown" {
		t.Fail()
	}
}

func TestNewInlineQueryResultArticleHTML(t *testing.T) {
	result := telegram.NewInlineQueryResultArticleHTML("id", "title", "<b>message</b>")

	if result.Type != "article" ||
		result.ID != "id" ||
		result.Title != "title" ||
		result.InputMessageContent.(telegram.InputTextMessageContent).Text != "<b>message</b>" ||
		result.InputMessageContent.(telegram.InputTextMessageContent).ParseMode != "HTML" {
		t.Fail()
	}
}

func TestNewInlineQueryResultGIF(t *testing.T) {
	result := telegram.NewInlineQueryResultGIF("id", "google.com")

	if result.Type != "gif" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultMPEG4GIF(t *testing.T) {
	result := telegram.NewInlineQueryResultMPEG4GIF("id", "google.com")

	if result.Type != "mpeg4_gif" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultPhoto(t *testing.T) {
	result := telegram.NewInlineQueryResultPhoto("id", "google.com")

	if result.Type != "photo" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultPhotoWithThumb(t *testing.T) {
	result := telegram.NewInlineQueryResultPhotoWithThumb("id", "google.com", "thumb.com")

	if result.Type != "photo" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.ThumbURL != "thumb.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultVideo(t *testing.T) {
	result := telegram.NewInlineQueryResultVideo("id", "google.com")

	if result.Type != "video" ||
		result.ID != "id" ||
		result.URL != "google.com" {
		t.Fail()
	}
}

func TestNewInlineQueryResultAudio(t *testing.T) {
	result := telegram.NewInlineQueryResultAudio("id", "google.com", "title")

	if result.Type != "audio" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.Title != "title" {
		t.Fail()
	}
}

func TestNewInlineQueryResultVoice(t *testing.T) {
	result := telegram.NewInlineQueryResultVoice("id", "google.com", "title")

	if result.Type != "voice" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.Title != "title" {
		t.Fail()
	}
}

func TestNewInlineQueryResultDocument(t *testing.T) {
	result := telegram.NewInlineQueryResultDocument("id", "google.com", "title", "mime/type")

	if result.Type != "document" ||
		result.ID != "id" ||
		result.URL != "google.com" ||
		result.Title != "title" ||
		result.MimeType != "mime/type" {
		t.Fail()
	}
}

func TestNewInlineQueryResultLocation(t *testing.T) {
	result := telegram.NewInlineQueryResultLocation("id", "name", 40, 50)

	if result.Type != "location" ||
		result.ID != "id" ||
		result.Title != "name" ||
		result.Latitude != 40 ||
		result.Longitude != 50 {
		t.Fail()
	}
}

func TestNewEditMessageText(t *testing.T) {
	edit := telegram.NewEditMessageText(ChatID, ReplyToMessageID, "new text")

	if edit.Text != "new text" ||
		edit.BaseEdit.ChatID != ChatID ||
		edit.BaseEdit.MessageID != ReplyToMessageID {
		t.Fail()
	}
}

func TestNewEditMessageCaption(t *testing.T) {
	edit := telegram.NewEditMessageCaption(ChatID, ReplyToMessageID, "new caption")

	if edit.Caption != "new caption" ||
		edit.BaseEdit.ChatID != ChatID ||
		edit.BaseEdit.MessageID != ReplyToMessageID {
		t.Fail()
	}
}

func TestNewEditMessageReplyMarkup(t *testing.T) {
	markup := telegram.InlineKeyboardMarkup{
		InlineKeyboard: [][]telegram.InlineKeyboardButton{
			[]telegram.InlineKeyboardButton{
				telegram.InlineKeyboardButton{Text: "test"},
			},
		},
	}

	edit := telegram.NewEditMessageReplyMarkup(ChatID, ReplyToMessageID, markup)

	if edit.ReplyMarkup.InlineKeyboard[0][0].Text != "test" ||
		edit.BaseEdit.ChatID != ChatID ||
		edit.BaseEdit.MessageID != ReplyToMessageID {
		t.Fail()
	}

}

func TestNewDice(t *testing.T) {
	dice := telegram.NewDice(42)

	if dice.ChatID != 42 ||
		dice.Emoji != "" {
		t.Fail()
	}
}

func TestNewDiceWithEmoji(t *testing.T) {
	dice := telegram.NewDiceWithEmoji(42, "🏀")

	if dice.ChatID != 42 ||
		dice.Emoji != "🏀" {
		t.Fail()
	}
}
