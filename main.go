package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ivahaev/amigo"
	"github.com/joho/godotenv"
)

// // Creating hanlder functions
func AgentCompleteHandler(a *amigo.Amigo, agentsQueue map[string]int64, m map[string]string) {
	fmt.Printf("AgentComplete event received: %v\n", m)
	go addAgentToPause(a, agentsQueue, m["MemberName"], m["Queue"], int64(time.Now().Unix()))
}

func DefaultHandler(m map[string]string) {
	fmt.Printf("Event received: %v\n", m)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	agentsQueue := make(map[string]int64)

	settings := &amigo.Settings{Username: os.Getenv("AMI_USERNAME"), Password: os.Getenv("AMI_PASSWORD"), Host: os.Getenv("AMI_HOST")}
	a := amigo.New(settings)

	a.Connect()

	queues := strings.Split(os.Getenv("QUEUESgive"), ",")
	fmt.Println(queues)

	// Listen for connection events
	a.On("connect", func(message string) {
		log.Println("Connected", message)
		unPauseAllAgents(a, queues)
	})
	a.On("error", func(message string) {
		fmt.Println("Connection error:", message)
	})

	// Registering handler function for event "DeviceStateChange"
	a.RegisterHandler("AgentComplete", func(m map[string]string) {
		AgentCompleteHandler(a, agentsQueue, m)
	})

	a.RegisterHandler("AsyncEvent", func(m map[string]string) {
		// Process the asynchronous response here
		fmt.Println("Received AsyncEvent:", m)
	})

	// goroutine for checking agents status in queue
	go checkAgentsStatus(a, agentsQueue)
	ch := make(chan bool)
	<-ch

}
