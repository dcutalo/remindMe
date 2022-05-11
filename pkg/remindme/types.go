package remindme

import "time"

type CreateReminder struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Tags         []string  `json:"tags"`
	ReminderTime time.Time `json:"reminderTime"`
	RemindType   string    `json:"reminderType"`
}
