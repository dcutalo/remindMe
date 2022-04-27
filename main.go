package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/mux"
)

// Setup HTTP server
// Create routes
// Create route handlers
// Create models for our data structures

// Setup temproal
// Create database
//

const (
	port = "8080"
)

type CreateReminder struct {
	Title        string    `json:"title"`
	Descrption   string    `json:"descreption"`
	Tags         []string  `json:"tags"`
	ReminderTime time.Time `json:"reminderTime"`
	RemindType   string    `json:"reminderType"`
}

var (
	BotToken = flag.String("token", "", "Bot token")
)

func main() {
	r := mux.NewRouter()
	// returns instance of discord bot
	discord, err := discordgo.New("Bot " + "")
	if err != nil {
		log.Fatalf("Failed to create discord bot %s", err)
	}
	discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})

	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %s", err)
	}
	defer discord.Close()

	sendMessage(discord)

	r.HandleFunc("/reminder", CreateReminderHandler).Methods("POST")
	r.HandleFunc("/reminder/{id}", UpdateReminderHandler).Methods("PUT")
	r.HandleFunc("/reminder/{id}", DeleteReminderHandler).Methods("DELETE")
	r.HandleFunc("/reminder/{id}", GetReminderHandler).Methods("GET")
	r.HandleFunc("/reminder", SearchReminderHandler).Methods("GET")

	log.Printf("Starting server on [%s]", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func sendMessage(discord *discordgo.Session) {
	message, err := discord.ChannelMessageSend(
		"724706617051840685", "Hello Ping Pong bot here!",
	)
	if err != nil {
		log.Printf("Failed to send message %s", err)
	}
	log.Printf("message to be sent: %s", message)
}

func CreateReminderHandler(w http.ResponseWriter, r *http.Request) {
	var reminder CreateReminder
	json.NewDecoder(r.Body).Decode(&reminder)

	fmt.Printf("%v\n", reminder)

	w.Write([]byte("create"))
	w.WriteHeader(200)
}

func UpdateReminderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update"))
	w.WriteHeader(200)
}

func DeleteReminderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
	w.WriteHeader(200)
}

func GetReminderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	w.Write([]byte("Get"))
	w.WriteHeader(200)
}

func SearchReminderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search"))
	w.WriteHeader(200)
}
