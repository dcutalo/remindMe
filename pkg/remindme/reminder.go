package remindme

import (
	"log"
	"time"
)

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

func (rm *ReminderManager) DeleteReminder(reminder DeleteReminder) error {
	sqlStatement := `
	DELETE FROM reminder WHERE remind_id=$1`

	res, err := rm.Db.Exec(sqlStatement, reminder.RemindID)
	if err != nil {
		log.Fatalf("Failed to delete from database table reminder: %s", err)
		return err
	}

	log.Printf("Result of delete query: %s", res)
	return nil
}

// not updating $3 and 4
func (rm *ReminderManager) UpdateReminder(reminder UpdateReminder) error {
	sqlStatement := `
	UPDATE reminder SET title = $2, reminder_message = $3, time_to_remind = $4
	WHERE remind_id=$1`

	res, err := rm.Db.Exec(sqlStatement, reminder.RemindID, reminder.Title, reminder.Description, reminder.ReminderTime)
	if err != nil {
		log.Fatalf("Failed to update from database table reminder: %s", err)
		return err
	}

	log.Printf("Result of update query: %s", res)
	log.Printf("reminder_message: %s", reminder.Description)
	log.Printf("time_to_remind: %s", reminder.ReminderTime)
	return nil
}

func (rm *ReminderManager) GetReminder(reminder GetReminder) error {
	sqlStatement := `
	SELECT * FROM reminder WHERE remind_id = $1`

	res, err := rm.Db.Exec(sqlStatement, reminder.RemindID)
	if err != nil {
		log.Fatalf("Failed to get from database table reminder: %s", err)
		return err
	}

	log.Printf("Result of get query: %s", res)
	return nil
}
