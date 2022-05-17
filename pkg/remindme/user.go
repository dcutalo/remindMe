package remindme

import (
	"log"
	"time"
)

func (rm *ReminderManager) CreateInsertUser(user CreateUser) error {
	sqlStatement := `
	INSERT INTO product_user (user_id, user_name, channel_id)
	VALUES ($1, $2, $3)`

	res, err := rm.Db.Exec(sqlStatement, time.Now().Unix(), user.UserName, user.ChannelId)
	if err != nil {
		log.Fatalf("Failed to insert into database table product_user: %s", err)
		return err
	}

	log.Printf("Result of insert query: %s", res)
	return nil
}

func (rm *ReminderManager) DeleteUser(user DeleteUser) error {
	sqlStatement := `
	DELETE FROM product_user WHERE user_id=$1`

	res, err := rm.Db.Exec(sqlStatement, user.UserID)
	if err != nil {
		log.Fatalf("Failed to delete from database table product_user: %s", err)
		return err
	}

	log.Printf("Result of delete query: %s", res)
	return nil
}
