package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// APIResponse is a response from the Telegram API with the result
// stored raw.
type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional
	RetryAfter      int   `json:"retry_after"`        // optional
}

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// Clear discards all unprocessed incoming updates.
func (ch UpdatesChannel) Clear() {
	for len(ch) != 0 {
		<-ch
	}
}

// User is a user on Telegram.
type User struct {
	ID                      int    `json:"id"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`                   // optional
	Username                string `json:"username"`                    // optional
	LanguageCode            string `json:"language_code"`               // optional
	IsBot                   bool   `json:"is_bot"`                      // optional
	CanJoinGroups           bool   `json:"can_join_groups"`             // optional. Returned only in getMe
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // optional. Returned only in getMe
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // optional. Returned only in getMe
}

// String displays a simple text version of a user.
//
// It is normally a user's username, but falls back to a first/last
// name as available.
func (u *User) String() string {
	if u == nil {
		return ""
	}
	if u.Username != "" {
		return u.Username
	}

	name := u.FirstName
	if u.LastName != "" {
		name += " " + u.LastName
	}

	return name
}

// GroupChat is a group chat.
type GroupChat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Chat contains information about the place a message was sent.
type Chat struct {
	ID               int64            `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`                 // optional
	UserName         string           `json:"username"`              // optional
	FirstName        string           `json:"first_name"`            // optional
	LastName         string           `json:"last_name"`             // optional
	Photo            *ChatPhoto       `json:"photo"`                 // optional
	Description      string           `json:"description,omitempty"` // optional
	InviteLink       string           `json:"invite_link,omitempty"` // optional
	PinnedMessage    *Message         `json:"pinned_message"`        // optional
	Permissions      *ChatPermissions `json:"permissions"`           // optional
	SlowModelDelay   int              `json:"slow_mode_delay"`       // optional
	StickerSetName   string           `json:"sticker_set_name"`      // optional
	CanSetStickerSet bool             `json:"can_set_sticker_set"`   // optional
}

// IsPrivate returns if the Chat is a private conversation.
func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsSuperGroup returns if the Chat is a supergroup.
func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

// IsChannel returns if the Chat is a channel.
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

// ChatConfig returns a ChatConfig struct for chat related methods.
func (c Chat) ChatConfig() ChatConfig {
	return ChatConfig{ChatID: c.ID}
}

// Message is returned by almost every request, and contains data about
// almost anything.
type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"from"` // optional
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from"`            // optional
	ForwardFromChat       *Chat                 `json:"forward_from_chat"`       // optional
	ForwardFromMessageID  int                   `json:"forward_from_message_id"` // optional
	ForwardSignature      string                `json:"forward_signature"`       // optional
	ForwardSenderName     string                `json:"forward_sender_name"`     // optional
	ForwardDate           int                   `json:"forward_date"`            // optional
	ReplyToMessage        *Message              `json:"reply_to_message"`        // optional
	ViaBot                *User                 `json:"via_bot"`                 // optional
	EditDate              int                   `json:"edit_date"`               // optional
	MediaGroupID          string                `json:"media_group_id"`          // optional
	AuthorSignature       string                `json:"author_signature"`        // optional
	Text                  string                `json:"text"`                    // optional
	Entities              *[]MessageEntity      `json:"entities"`                // optional
	Animation             *Animation            `json:"animation"`               // optional
	Audio                 *Audio                `json:"audio"`                   // optional
	Document              *Document             `json:"document"`                // optional
	Photo                 *[]PhotoSize          `json:"photo"`                   // optional
	Sticker               *Sticker              `json:"sticker"`                 // optional
	Video                 *Video                `json:"video"`                   // optional
	VideoNote             *VideoNote            `json:"video_note"`              // optional
	Voice                 *Voice                `json:"voice"`                   // optional
	Caption               string                `json:"caption"`                 // optional
	CaptionEntities       *[]MessageEntity      `json:"caption_entities"`        // optional
	Contact               *Contact              `json:"contact"`                 // optional
	Dice                  *Dice                 `json:"dice"`                    // optional
	Game                  *Game                 `json:"game"`                    // optional
	Poll                  *Poll                 `json:"poll"`                    // optional // Not implemented yet
	Venue                 *Venue                `json:"venue"`                   // optional
	Location              *Location             `json:"location"`                // optional
	NewChatMembers        *[]User               `json:"new_chat_members"`        // optional
	LeftChatMember        *User                 `json:"left_chat_member"`        // optional
	NewChatTitle          string                `json:"new_chat_title"`          // optional
	NewChatPhoto          *[]PhotoSize          `json:"new_chat_photo"`          // optional
	DeleteChatPhoto       bool                  `json:"delete_chat_photo"`       // optional
	GroupChatCreated      bool                  `json:"group_chat_created"`      // optional
	SuperGroupChatCreated bool                  `json:"supergroup_chat_created"` // optional
	ChannelChatCreated    bool                  `json:"channel_chat_created"`    // optional
	MigrateToChatID       int64                 `json:"migrate_to_chat_id"`      // optional
	MigrateFromChatID     int64                 `json:"migrate_from_chat_id"`    // optional
	PinnedMessage         *Message              `json:"pinned_message"`          // optional
	Invoice               *Invoice              `json:"invoice"`                 // optional
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment"`      // optional
	ConnectedWebsite      string                `json:"connected_website"`       // optional
	PassportData          *PassportData         `json:"passport_data,omitempty"` // optional
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`            // optional
}

// Time converts the message timestamp into a Time.
func (m *Message) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(*m.Entities) == 0 {
		return false
	}

	entity := (*m.Entities)[0]
	return entity.Offset == 0 && entity.IsCommand()
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is removed. Use
// CommandWithAt() if you do not want that.
func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandWithAt checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is not removed. Use Command()
// if you want that.
func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	return m.Text[1:entity.Length]
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
//
// Note: The first character after the command name is omitted:
// - "/foo bar baz" yields "bar baz", not " bar baz"
// - "/foo-bar baz" yields "bar baz", too
// Even though the latter is not a command conforming to the spec, the API
// marks "/foo" as command entity.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := (*m.Entities)[0]
	if len(m.Text) == entity.Length {
		return "" // The command makes up the whole message
	}

	return m.Text[entity.Length+1:]
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url"`      // optional
	User     *User  `json:"user"`     // optional
	Language string `json:"language"` // optional
}

// ParseURL attempts to parse a URL contained within a MessageEntity.
func (e MessageEntity) ParseURL() (*url.URL, error) {
	if e.URL == "" {
		return nil, errors.New(ErrBadURL)
	}

	return url.Parse(e.URL)
}

// IsMention returns true if the type of the message entity is "mention" (@username).
func (e MessageEntity) IsMention() bool {
	return e.Type == "mention"
}

// IsHashtag returns true if the type of the message entity is "hashtag".
func (e MessageEntity) IsHashtag() bool {
	return e.Type == "hashtag"
}

// IsCashtag returns true if the type of the message entity is "cashtag".
func (e MessageEntity) IsCashtag() bool {
	return e.Type == "cashtag"
}

// IsCommand returns true if the type of the message entity is "bot_command".
func (e MessageEntity) IsCommand() bool {
	return e.Type == "bot_command"
}

// IsURL returns true if the type of the message entity is "url".
func (e MessageEntity) IsURL() bool {
	return e.Type == "url"
}

// IsEmail returns true if the type of the message entity is "email".
func (e MessageEntity) IsEmail() bool {
	return e.Type == "email"
}

// IsPhoneNumber returns true if the type of the message entity is "phone_number".
func (e MessageEntity) IsPhoneNumber() bool {
	return e.Type == "phone_number"
}

// IsBold returns true if the type of the message entity is "bold" (bold text).
func (e MessageEntity) IsBold() bool {
	return e.Type == "bold"
}

// IsItalic returns true if the type of the message entity is "italic" (italic text).
func (e MessageEntity) IsItalic() bool {
	return e.Type == "italic"
}

// IsCode returns true if the type of the message entity is "code" (monowidth string).
func (e MessageEntity) IsCode() bool {
	return e.Type == "code"
}

// IsPre returns true if the type of the message entity is "pre" (monowidth block).
func (e MessageEntity) IsPre() bool {
	return e.Type == "pre"
}

// IsTextLink returns true if the type of the message entity is "text_link" (clickable text URL).
func (e MessageEntity) IsTextLink() bool {
	return e.Type == "text_link"
}

// IsTextMention returns true if the type of the message entity is "text_mention".
func (e MessageEntity) IsTextMention() bool {
	return e.Type == "text_mention"
}

// PhotoSize contains information about photos.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"` // optional
}

// Animation contains information about an animation.
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`     // optional
	FileName     string     `json:"file_name"` // optional
	MimeType     string     `json:"mime_type"` // optional
	FileSize     int        `json:"file_size"` // optional
}

// Audio contains information about audio.
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer"` // optional
	Title        string     `json:"title"`     // optional
	MimeType     string     `json:"mime_type"` // optional
	FileSize     int        `json:"file_size"` // optional
	Thumbnail    *PhotoSize `json:"thumb"`     // optional
}

// Document contains information about a document.
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumb"`     // optional
	FileName     string     `json:"file_name"` // optional
	MimeType     string     `json:"mime_type"` // optional
	FileSize     int        `json:"file_size"` // optional
}

// Video contains information about a video.
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumb"`     // optional
	MimeType     string     `json:"mime_type"` // optional
	FileSize     int        `json:"file_size"` // optional
}

// VideoNote contains information about a video.
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumb"`     // optional
	FileSize     int        `json:"file_size"` // optional
}

// Voice contains information about a voice.
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"` // optional
	FileSize     int    `json:"file_size"` // optional
}

// Contact contains information about a contact.
//
// Note that LastName and UserID may be empty.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"` // optional
	UserID      int    `json:"user_id"`   // optional
	Vcard       string `json:"vcard"`     //optional
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text"`
	VoterCount string `json:"voter_count"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               *[]PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"json"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id"`
	Explanation           string           `json:"explanation"`
	ExplanationEntities   *[]MessageEntity `json:"explanation_entities"`
	OpenPeriod            int              `json:"open_period"`
	CloseDate             int              `json:"close_date"`
}

// Location contains information about a place.
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue contains information about a venue, including its Location.
type Venue struct {
	Location       Location `json:"location"`
	Title          string   `json:"title"`
	Address        string   `json:"address"`
	FoursquareID   string   `json:"foursquare_id"`   // optional
	FoursquareType string   `json:"foursquare_type"` // optional
}

// UserProfilePhotos contains a set of user profile photos.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File contains information about a file to download from Telegram.
type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"` // optional
	FilePath     string `json:"file_path"` // optional
}

// Link returns a full path to the download URL for a File.
//
// It requires the Bot Token to create the link.
func (f *File) Link(token string) string {
	return fmt.Sprintf(FileEndpoint, token, f.FilePath)
}

// ReplyKeyboardMarkup allows the Bot to set a custom keyboard.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`   // optional
	OneTimeKeyboard bool               `json:"one_time_keyboard"` // optional
	Selective       bool               `json:"selective"`         // optional
}

// KeyboardButton is a button within a custom keyboard.
type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact"`
	RequestLocation bool                    `json:"request_location"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`
}

// KeyboardButtonPollType represents type of a poll
type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

// ReplyKeyboardRemove allows the Bot to hide a custom keyboard.
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

// InlineKeyboardMarkup is a custom keyboard presented for an inline bot.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is a button within a custom keyboard for
// inline query responses.
//
// Note that some values are references as even an empty string
// will change behavior.
//
// CallbackGame, if set, MUST be first button in first row.
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	URL                          *string       `json:"url,omitempty"` // optional
	LoginURL                     *LoginURL     `json:"login_url"`
	CallbackData                 *string       `json:"callback_data,omitempty"`                    // optional
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`              // optional
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"` // optional
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`                    // optional
	Pay                          bool          `json:"pay,omitempty"`                              // optional
}

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user
type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

// CallbackQuery is data sent when a keyboard button with callback data
// is clicked.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`           // optional
	InlineMessageID string   `json:"inline_message_id"` // optional
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`            // optional
	GameShortName   string   `json:"game_short_name"` // optional
}

// ForceReply allows the Bot to have users directly reply to it without
// additional interaction.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective"` // optional
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatMember is information about a member in a chat.
type ChatMember struct {
	User                  *User  `json:"user"`
	Status                string `json:"status"`
	CustomTitle           string `json:"custom_title,omitempty"`              // optional
	UntilDate             int64  `json:"until_date,omitempty"`                // optional
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`             // optional
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`         // optional
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`         // optional
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`       // optional
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`      // optional
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`       // optional
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`           // optional
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`          // optional
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`          // optional
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`         // optional
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`   // optional
	CanSendPolls          bool   `json:"can_send_polls"`                      // optional
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`   // optional
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"` // optional
}

// IsCreator returns if the ChatMember was the creator of the chat.
func (chat ChatMember) IsCreator() bool { return chat.Status == "creator" }

// IsAdministrator returns if the ChatMember is a chat administrator.
func (chat ChatMember) IsAdministrator() bool { return chat.Status == "administrator" }

// IsMember returns if the ChatMember is a current member of the chat.
func (chat ChatMember) IsMember() bool { return chat.Status == "member" }

// HasLeft returns if the ChatMember left the chat.
func (chat ChatMember) HasLeft() bool { return chat.Status == "left" }

// WasKicked returns if the ChatMember was kicked from the chat.
func (chat ChatMember) WasKicked() bool { return chat.Status == "kicked" }

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

// BotCommand represents a bot command.
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// InputMediaPhoto contains a photo for displaying as part of a media group.
type InputMediaPhoto struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

// InputMediaVideo contains a video for displaying as part of a media group.
type InputMediaVideo struct {
	Type  string `json:"type"`
	Media string `json:"media"`
	// thumb intentionally missing as it is not currently compatible
	Caption           string `json:"caption"`
	ParseMode         string `json:"parse_mode"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
	Duration          int    `json:"duration"`
	SupportsStreaming bool   `json:"supports_streaming"`
}

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct{}

// InputMediaAudio represents an audio file to be treated as music to be sent.
type InputMediaAudio struct{}

// InputMediaDocument represents a general file to be sent.
type InputMediaDocument struct{}

// Sticker contains information about a sticker.
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb"`         // optional
	Emoji        string        `json:"emoji"`         // optional
	SetName      string        `json:"set_name"`      // optional
	MaskPosition *MaskPosition `json:"mask_position"` // optional
	FileSize     int           `json:"file_size"`     // optional
}

// StickerSet contains information about an sticker set.
type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	IsAnimated    bool      `json:"is_animated"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}

// MaskPosition describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

// Game is a game within Telegram.
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

// CallbackGame is for starting a game in an inline keyboard button.
type CallbackGame struct{}

// GameHighScore is a user's score and position on the leaderboard.
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

// WebhookInfo is information about a currently set webhook.
type WebhookInfo struct {
	URL                  string `json:"url"`
	HasCustomCertificate bool   `json:"has_custom_certificate"`
	PendingUpdateCount   int    `json:"pending_update_count"`
	LastErrorDate        int    `json:"last_error_date"`    // optional
	LastErrorMessage     string `json:"last_error_message"` // optional
	MaxConnections       int    `json:"max_connections"`    // optional
}

// IsSet returns true if a webhook is currently set.
func (info WebhookInfo) IsSet() bool {
	return info.URL != ""
}

// Error is an error containing extra information returned by the Telegram API.
type Error struct {
	Code    int
	Message string
	ResponseParameters
}

func (e Error) Error() string {
	return e.Message
}
