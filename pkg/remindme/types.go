package remindme

import "time"

type Reminder struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ReminderId struct {
	RemindID int `json:"remind_id"`
}

type CreateReminder struct {
	Reminder
	Tags         []string  `json:"tags"`
	ReminderTime time.Time `json:"reminderTime"`
	RemindType   string    `json:"reminderType"`
}

type UpdateReminder struct {
	Reminder
	ReminderId
	ReminderTime time.Time `json:"reminderTime"`
}

type DeleteReminder struct {
	ReminderId
}

type GetReminder struct {
	ReminderId
}

// user structs
type User struct {
	UserId
	UserName  string `json:"user_name"`
	ChannelId string `json:"channel_id"`
}

type UserId struct {
	UserID int `json:"user_id"`
}

type CreateUser struct {
	User
}

type DeleteUser struct {
	UserId
}
