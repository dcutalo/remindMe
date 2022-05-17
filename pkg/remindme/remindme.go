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

func (rmapi *RemindMeAPI) DeleteReminderHandler(w http.ResponseWriter, r *http.Request) {
	var reminderID DeleteReminder
	if err := json.NewDecoder(r.Body).Decode(&reminderID); err != nil {
		w.Write([]byte("unable to decode body"))
		w.WriteHeader(500)
		return
	}

	if err := rmapi.ReminderManager.DeleteReminder(reminderID); err != nil {
		w.Write([]byte(fmt.Sprintf("unable to delete reminder %s", err.Error())))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("reminder deleted"))
	w.WriteHeader(200)
}

func (rmapi *RemindMeAPI) UpdateReminderHandler(w http.ResponseWriter, r *http.Request) {
	var reminder UpdateReminder
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		w.Write([]byte("unable to decode body"))
		w.WriteHeader(500)
		return
	}

	if err := rmapi.ReminderManager.UpdateReminder(reminder); err != nil {
		w.Write([]byte(fmt.Sprintf("unable to update reminder %s", err.Error())))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("update reminder"))
	w.WriteHeader(200)
}

func (rmapi *RemindMeAPI) GetReminderHandler(w http.ResponseWriter, r *http.Request) {
	var reminder GetReminder
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		w.Write([]byte("unable to decode body"))
		w.WriteHeader(500)
		return
	}

	if err := rmapi.ReminderManager.GetReminder(reminder); err != nil {
		w.Write([]byte(fmt.Sprintf("unable to get reminder %s", err.Error())))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("Get reminder"))
	w.WriteHeader(200)
}

// user handlers
func (rmapi *RemindMeAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user CreateUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Write([]byte("unable to decode body"))
		w.WriteHeader(500)
		return
	}

	if err := rmapi.ReminderManager.CreateInsertUser(user); err != nil {
		w.Write([]byte(fmt.Sprintf("unable to create user %s", err.Error())))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("created user"))
	w.WriteHeader(200)
}

func (rmapi *RemindMeAPI) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user DeleteUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Write([]byte("unable to decode body"))
		w.WriteHeader(500)
		return
	}

	if err := rmapi.ReminderManager.DeleteUser(user); err != nil {
		w.Write([]byte(fmt.Sprintf("unable to delete user %s", err.Error())))
		w.WriteHeader(500)
		return
	}

	w.Write([]byte("delete user"))
	w.WriteHeader(200)
}
