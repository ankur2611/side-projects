package dto

type CommandsRequest struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message"`
}

type WebHookResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

type TgResponse struct {
	Ok          bool    `json:"ok"`
	Result      *Result `json:"result"`
	ErrorCode   int     `json:"error_code"`
	Description string  `json:"description"`
}

type Result struct {
	MessageID int        `json:"message_id"`
	From      *User      `json:"from"`
	Chat      *Chat      `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Animation *Animation `json:"animation"`
	Document  *Document  `json:"document"`
}

type Animation struct {
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	FileSize     int        `json:"file_size"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	ThumbNail    *ThumbNail `json:"thumbnail"`
	Thumb        *Thumb     `json:"thumb"`
}

type ThumbNail struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type Thumb struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type Document struct {
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	FileSize     int        `json:"file_size"`
	ThumbNail    *ThumbNail `json:"thumbnail"`
	Thumb        *Thumb     `json:"thumb"`
}

type Message struct {
	Text           string `json:"text"`
	Chat           *Chat  `json:"chat"`
	NewChatMembers []User `json:"new_chat_members"`
	MessageID      int    `json:"message_id"`
	From           *User  `json:"from"`
	Date           int    `json:"date"`
	LeftChatMember *User  `json:"left_chat_member,omitempty"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type User struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}