package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	webhook "github.com/koron/go-github-webhook"
)

var command = flag.String("-c", "./trigger.sh", "command to execute for webhook")

func main() {
	flag.Parse()
	log.Println("starting init trigger...")
	exec(*command)

	log.Println("starting service...")

	// To verify webhook's payload, set secret by SetSecret().
	webhook.SetSecret([]byte(""))

	// Add a HandlerFunc to process webhook.
	http.HandleFunc("/", webhook.HandlePush(handler))

	// Start web server.
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handler(ev *webhook.Event) {
	push := ev.PushEvent()
	if push == nil {
		log.Printf("got empty push event\n")
		pretty("event: ", ev)
		return
	}
	pretty("push", push)
	log.Printf("starting trigger command: %v\n", *command)
	exec(*command)
	log.Printf("end of trigger command: %v\n", *command)
}

func pretty(key string, a interface{}) {
	b, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(key, string(b))
}
