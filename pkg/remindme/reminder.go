package remindme

import (
	"database/sql"
	"log"
	"time"
)

type ReminderManager struct {
	Db *sql.DB
}

// testing raw input into db
func (rm *ReminderManager) CreateInsertReminder(reminder CreateReminder) error {
	sqlStatement := `
	INSERT INTO reminder (remind_id, title, reminder_message, time_to_remind)
	VALUES ($1, $2, $3, $4)`

	res, err := rm.Db.Exec(sqlStatement, time.Now().Unix(), reminder.Title, reminder.Description, reminder.ReminderTime)
	if err != nil {
		log.Fatalf("Failed to insert into database table reminder: %s", err)
		return err
	}

	log.Printf("Result of insert query: %s", res)
	return nil
}
