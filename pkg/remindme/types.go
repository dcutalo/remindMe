package remindme

import (
	"time"
)

type Reminder struct {
	ReminderId
	Title        string    `json:"title"`
	Description  string    `json:"reminder_message"`
	ReminderTime time.Time `json:"time_to_remind"`
}

type ReminderId struct {
	RemindID int `json:"remind_id"`
}

type CreateReminder struct {
	Reminder
}

//Tags         []string  `json:"tags"`
//RemindType   string    `json:"reminderType"`

type UpdateReminder struct {
	Reminder
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

type UpdateUser struct {
	User
}

type DeleteUser struct {
	UserId
}

type GetUser struct {
	UserId
}
