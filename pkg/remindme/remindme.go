package remindme

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RemindMeAPI struct {
	ReminderManager ReminderManager
}

func (rmapi *RemindMeAPI) CreateReminderHandler(w http.ResponseWriter, r *http.Request) {
	var reminder CreateReminder
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		w.Write([]byte("unable to decode body"))
		w.WriteHeader(500)
		return
	}

	if err := rmapi.ReminderManager.CreateInsertReminder(reminder); err != nil {
		w.Write([]byte(fmt.Sprintf("unable to create reminder %s", err.Error())))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("created reminder"))
	w.WriteHeader(200)
}
