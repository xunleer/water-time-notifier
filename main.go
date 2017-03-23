package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"os"

	"fmt"

	"sync"

	gmux "github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type netNotifications struct {
	Notifications []notification `json:"notifications,omitempty"`
}

func processNetNotification(w http.ResponseWriter, req *http.Request) {
	var notifiers netNotifications
	fmt.Println(req.Body)
	err := json.NewDecoder(req.Body).Decode(&notifiers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, notifier := range notifiers.Notifications {
		notifier.Notice()
	}
}

func handleNetRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	mux := gmux.NewRouter()
	mux.HandleFunc("/notifications", processNetNotification).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(mux)

	var srvPort string
	if srvPort = os.Getenv("SRVPORT"); srvPort == "" {
		srvPort = "8080"
	}
	n.Run(":" + srvPort)
}

var wtNotice bool

func init() {
	flag.BoolVar(&wtNotice, "b", false, "Water time notice for each hour.")
}

func main() {
	flag.Parse()

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go handleNetRequest(&waitGroup)

	if wtNotice == true {
		timer := time.NewTimer(time.Hour * 1)
		notifier := notification{
			Application: "Water Time Notifier",
			Title:       "Water Time Now!",
			Message:     "Hey, you have worked for an hour, it's better to have some water now.",
		}

		for {
			<-timer.C
			err := notifier.Notice()
			if err != nil {
				log.Fatalln(err)
			}

			timer = time.NewTimer(time.Hour * 1)
		}
	}

	waitGroup.Wait()
}
