package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "io/ioutil"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("slack-auto-away", "Set a slack status")
	var mySelector *string = parser.Selector("s", "status", []string{"auto", "away"}, &argparse.Options{
		Required: true,
		Help:     "Set a slack status",
		Default:  "auto"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	SLACK_API := os.Getenv("SLACK_API")
	SLACK_SET_STATUS := "https://slack.com/api/users.setPresence?presence=" + *mySelector

	req, err := http.NewRequest("POST", SLACK_SET_STATUS, nil)
	req.Header.Add("Authorization", "Bearer " + SLACK_API)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	// f,_ := ioutil.ReadAll(resp.Body)
	// fmt.Print(string(f))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
