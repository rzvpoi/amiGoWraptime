package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ivahaev/amigo"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func checkAgentsStatus(a *amigo.Amigo, agentsQueue map[string]int64) {
	wraptime, err := strconv.Atoi(os.Getenv("WRAPTIME"))
	if err != nil {
		log.Panicln("Error:", err)
	}
	for {
		for agent, pause_time := range agentsQueue {
			if time.Now().Unix()-pause_time > int64(wraptime) {
				removeAgentFromPause(a, agentsQueue, agent)
			}

		}
		time.Sleep(100 * time.Millisecond)
	}
}

func addAgentToPause(a *amigo.Amigo, agentsQueue map[string]int64, memberName string, queue string, pause_time int64) {
	// extract the extension from Member Name (ex: 300 Fname Lname)
	exten := memberName[:3]

	if _, err := os.Stat("/.dockerenv"); err != nil {
		// if not running in docker cotainer read from .env file
		err = godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	allowedQueues := strings.Split(os.Getenv("QUEUES"), ",")
	var isQueueAllowed bool = false
	for _, allowedQueue := range allowedQueues {
		if strings.Compare(queue, allowedQueue) == 0 {
			isQueueAllowed = true
		}
	}
	if !isQueueAllowed {
		log.Println("Queue is not allowed!")
		return
	}

	if agentsQueue[exten] != 0 {
		log.Println("Error:", fmt.Errorf("Field `"+exten+"` already exists"))
		return
	}

	if a.Connected() {
		result, err := a.Action(map[string]string{
			"Action":    "QueuePause",
			"Interface": fmt.Sprintf("Local/%s@from-queue/n", exten),
			"Paused":    "true",
			"Reason":    "Wraptime pause",
			"ActionID":  "pause",
		})
		if err != nil {
			log.Println(result, err)
		}
	} else {
		log.Println("Error while trying to send command: Not connected to AMI!")
		return
	}

	agentsQueue[exten] = pause_time

	log.Printf("Agent %s added to pause for 30 seconds", memberName)
}

func removeAgentFromPause(a *amigo.Amigo, agentsQueue map[string]int64, exten string) {
	if agentsQueue[exten] == 0 {
		log.Println("Error:", fmt.Errorf("Field `"+exten+"` already exists"))
		return
	}

	if a.Connected() {
		result, err := a.Action(map[string]string{
			"Action":    "QueuePause",
			"Interface": fmt.Sprintf("Local/%s@from-queue/n", exten),
			"Paused":    "false",
			"Reason":    "Wraptime unpause",
			"ActionID":  "unpause",
		})
		if err != nil {
			log.Println(result["Response"], err)
		}
	} else {
		log.Println("Error while trying to send command: Not connected to AMI!")
		return
	}

	delete(agentsQueue, exten)

	log.Printf("Agent %s removed from pause", exten)
}

func unPauseAllAgents(a *amigo.Amigo, queues []string) {
	for _, queue := range queues {
		if a.Connected() {
			result, err := a.Action(map[string]string{
				"Action":   "Command",
				"Command":  fmt.Sprintf("queue show %s", queue),
				"ActionID": "get queue status",
			})
			if err != nil {
				log.Println(result, err)
			}
			response := strings.Split(result["CommandResponse"], " ")
			for idx := range response {
				if (strings.Contains(response[idx], "Local/")) && (strings.Contains(response[idx+6], "(paused)")) {
					ext := response[idx][len("Local/")+1 : len("Local/")+4]
					result, err := a.Action(map[string]string{
						"Action":    "QueuePause",
						"Interface": fmt.Sprintf("Local/%s@from-queue/n", ext),
						"Paused":    "false",
						"Reason":    "Wraptime unpause",
						"ActionID":  "unpause",
					})
					if err != nil {
						log.Println(result["Response"], err)
					}
					log.Printf("Agent %s removed from pause before starting listening for events", ext)

				}
			}
		}
	}
}
