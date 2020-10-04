package telegram_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	telegram "github.com/abhaikollara/telegram-bot-api"
)

const (
	TestToken               = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"
	ChatID                  = 76918703
	SupergroupChatID        = -1001120141283
	ReplyToMessageID        = 35
	ExistingPhotoFileID     = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	ExistingDocumentFileID  = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	ExistingAudioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingVideoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	ExistingStickerFileID   = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

func getBot(t *testing.T) (*telegram.BotAPI, error) {
	bot, err := telegram.NewBotAPI(TestToken)
	bot.Debug = true

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	return bot, err
}

func TestNewBotAPI_notoken(t *testing.T) {
	_, err := telegram.NewBotAPI("")

	if err == nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetUpdates(t *testing.T) {
	bot, _ := getBot(t)

	u := telegram.NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessageReply(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ReplyToMessageID = ReplyToMessageID
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMessageForward(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewForward(ChatID, ChatID, ReplyToMessageID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhoto(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewPhotoUpload(ChatID, "tests/image.jpg")
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileBytes(t *testing.T) {
	bot, _ := getBot(t)

	data, _ := ioutil.ReadFile("tests/image.jpg")
	b := telegram.FileBytes{Name: "image.jpg", Bytes: data}

	msg := telegram.NewPhotoUpload(ChatID, b)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileReader(t *testing.T) {
	bot, _ := getBot(t)

	f, _ := os.Open("tests/image.jpg")
	reader := telegram.FileReader{Name: "image.jpg", Reader: f, Size: -1}

	msg := telegram.NewPhotoUpload(ChatID, reader)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewPhotoReply(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewPhotoUpload(ChatID, "tests/image.jpg")
	msg.ReplyToMessageID = ReplyToMessageID

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingPhoto(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewPhotoShare(ChatID, ExistingPhotoFileID)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewDocument(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewDocumentUpload(ChatID, "tests/image.jpg")
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingDocument(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewDocumentShare(ChatID, ExistingDocumentFileID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewAudio(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewAudioUpload(ChatID, "tests/audio.mp3")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"
	msg.MimeType = "audio/mpeg"
	msg.FileSize = 688
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingAudio(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewAudioShare(ChatID, ExistingAudioFileID)
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVoice(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewVoiceUpload(ChatID, "tests/voice.ogg")
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVoice(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewVoiceShare(ChatID, ExistingVoiceFileID)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithContact(t *testing.T) {
	bot, _ := getBot(t)

	contact := telegram.NewContact(ChatID, "5551234567", "Test")

	if _, err := bot.Send(contact); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithLocation(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(telegram.NewLocation(ChatID, 40, 40))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithVenue(t *testing.T) {
	bot, _ := getBot(t)

	venue := telegram.NewVenue(ChatID, "A Test Location", "123 Test Street", 40, 40)

	if _, err := bot.Send(venue); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVideo(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewVideoUpload(ChatID, "tests/video.mp4")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVideo(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewVideoShare(ChatID, ExistingVideoFileID)
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewVideoNote(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewVideoNoteUpload(ChatID, 240, "tests/videonote.mp4")
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingVideoNote(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewVideoNoteShare(ChatID, 240, ExistingVideoNoteFileID)
	msg.Duration = 10

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewSticker(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewStickerUpload(ChatID, "tests/image.jpg")

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingSticker(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewStickerShare(ChatID, ExistingStickerFileID)

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithNewStickerAndKeyboardHide(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewStickerUpload(ChatID, "tests/image.jpg")
	msg.ReplyMarkup = telegram.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      false,
	}
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithExistingStickerAndKeyboardHide(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewStickerShare(ChatID, ExistingStickerFileID)
	msg.ReplyMarkup = telegram.ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      false,
	}

	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithDice(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewDice(ChatID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

}

func TestSendWithDiceWithEmoji(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewDiceWithEmoji(ChatID, "🏀")
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

}

func TestGetFile(t *testing.T) {
	bot, _ := getBot(t)

	file := telegram.FileConfig{FileID: ExistingPhotoFileID}

	_, err := bot.GetFile(file)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendChatConfig(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.Send(telegram.NewChatAction(ChatID, telegram.ChatTyping))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendEditMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg, err := bot.Send(telegram.NewMessage(ChatID, "Testing editing."))
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	edit := telegram.EditMessageTextConfig{
		BaseEdit: telegram.BaseEdit{
			ChatID:    ChatID,
			MessageID: msg.MessageID,
		},
		Text: "Updated text.",
	}

	_, err = bot.Send(edit)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	bot, _ := getBot(t)

	_, err := bot.GetUserProfilePhotos(telegram.NewUserProfilePhotos(ChatID))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSetWebhookWithCert(t *testing.T) {
	bot, _ := getBot(t)

	time.Sleep(time.Second * 2)

	bot.RemoveWebhook()

	wh := telegram.NewWebhookWithCert("https://example.com/telegram-test/"+bot.Token, "tests/cert.pem")
	_, err := bot.SetWebhook(wh)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	_, err = bot.GetWebhookInfo()
	if err != nil {
		t.Error(err)
	}
	bot.RemoveWebhook()
}

func TestSetWebhookWithoutCert(t *testing.T) {
	bot, _ := getBot(t)

	time.Sleep(time.Second * 2)

	bot.RemoveWebhook()

	wh := telegram.NewWebhook("https://example.com/telegram-test/" + bot.Token)
	_, err := bot.SetWebhook(wh)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		t.Error(err)
	}
	if info.MaxConnections == 0 {
		t.Errorf("Expected maximum connections to be greater than 0")
	}
	if info.LastErrorDate != 0 {
		t.Errorf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	bot.RemoveWebhook()
}

func TestUpdatesChan(t *testing.T) {
	bot, _ := getBot(t)

	var ucfg telegram.UpdateConfig = telegram.NewUpdate(0)
	ucfg.Timeout = 60
	_, err := bot.GetUpdatesChan(ucfg)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSendWithMediaGroup(t *testing.T) {
	bot, _ := getBot(t)

	cfg := telegram.NewMediaGroup(ChatID, []interface{}{
		telegram.NewInputMediaPhoto("https://i.imgur.com/unQLJIb.jpg"),
		telegram.NewInputMediaPhoto("https://i.imgur.com/J5qweNZ.jpg"),
		telegram.NewInputMediaVideo("https://i.imgur.com/F6RmI24.mp4"),
	})
	_, err := bot.Send(cfg)
	if err != nil {
		t.Error(err)
	}
}

func ExampleNewBotAPI() {
	bot, err := telegram.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Username)

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)

		msg := telegram.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

func ExampleNewWebhook() {
	bot, err := telegram.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Username)

	_, err = bot.SetWebhook(telegram.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
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
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func ExampleWebhookHandler() {
	bot, err := telegram.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Username)

	_, err = bot.SetWebhook(telegram.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
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

	http.HandleFunc("/"+bot.Token, func(w http.ResponseWriter, r *http.Request) {
		update, err := bot.HandleUpdate(r)
		if err != nil {
			log.Printf("%+v\n", err.Error())
		} else {
			log.Printf("%+v\n", *update)
		}
	})

	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
}

func ExampleAnswerInlineQuery() {
	bot, err := telegram.NewBotAPI("MyAwesomeBotToken") // create new bot
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.Username)

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery == nil { // if no inline query, ignore it
			continue
		}

		article := telegram.NewInlineQueryResultArticle(update.InlineQuery.ID, "Echo", update.InlineQuery.Query)
		article.Description = update.InlineQuery.Query

		inlineConf := telegram.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results:       []interface{}{article},
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}

func TestDeleteMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	message, _ := bot.Send(msg)

	deleteMessageConfig := telegram.DeleteMessageConfig{
		ChatID:    message.Chat.ID,
		MessageID: message.MessageID,
	}
	_, err := bot.DeleteMessage(deleteMessageConfig)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestPinChatMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewMessage(SupergroupChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	message, _ := bot.Send(msg)

	pinChatMessageConfig := telegram.PinChatMessageConfig{
		ChatID:              message.Chat.ID,
		MessageID:           message.MessageID,
		DisableNotification: false,
	}
	_, err := bot.PinChatMessage(pinChatMessageConfig)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestUnpinChatMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := telegram.NewMessage(SupergroupChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	message, _ := bot.Send(msg)

	// We need pin message to unpin something
	pinChatMessageConfig := telegram.PinChatMessageConfig{
		ChatID:              message.Chat.ID,
		MessageID:           message.MessageID,
		DisableNotification: false,
	}
	_, err := bot.PinChatMessage(pinChatMessageConfig)

	unpinChatMessageConfig := telegram.UnpinChatMessageConfig{
		ChatID: message.Chat.ID,
	}
	_, err = bot.UnpinChatMessage(unpinChatMessageConfig)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
