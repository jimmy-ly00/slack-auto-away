package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("slack-auto-away", "Set a slack status")
	var mySelector *string = parser.Selector("s", "status", []string{"active", "auto"}, &argparse.Options{
		Required: true,
		Help:     "Set a slack status",
		Default:  "auto"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	//CHANGE ME
	SLACK_URL := "https://[TEAMNAME].slack.com/"
	API_TOKEN := "xoxs-.............."
	SLACK_SET_STATUS := SLACK_URL + "/api/users.setPresence?presence=" + *mySelector + "&token=" + API_TOKEN

	res, err := http.Get(SLACK_SET_STATUS)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
}
