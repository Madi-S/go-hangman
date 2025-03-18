package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const DEFAULT_WORD = "golang"
const URL = "https://random-words5.p.rapidapi.com/getMultipleRandom?count=1"

func getRandomWord(client *http.Client) string {
	req, err := http.NewRequest("GET", URL, bytes.NewReader([]byte{}))
	if err != nil {
		return DEFAULT_WORD
	}
	// Visit https://rapidapi.com/sheharyar566/api/random-words5 for your api key
	req.Header.Set("x-rapidapi-key", "???")
	req.Header.Set("x-rapidapi-host", "random-words5.p.rapidapi.com")

	resp, err := client.Do(req)
	if err != nil {
		return DEFAULT_WORD
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return DEFAULT_WORD
	}

	var data []string
	err = json.Unmarshal(body, &data)
	if err != nil {
		return DEFAULT_WORD
	}

	return data[0]
}
