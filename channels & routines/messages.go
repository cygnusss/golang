package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DadJokesResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func GetMessages() {
	go func() {
		for {
			var msg DadJokesResponse
			c := &http.Client{}

			req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
			if err != nil {
				fmt.Printf("Error at creating a get request: %s", err.Error())
			}

			req.Header.Add("Accept", "application/json")
			resp, err := c.Do(req)
			if err != nil {
				fmt.Printf("Error at sending a get request: %s", err.Error())
			}

			if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
				fmt.Printf("Error at decoding a response: %s", err.Error())
			}

			JobQueue <- Job{msg.Joke}
		}
	}()
}
