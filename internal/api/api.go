package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiKeyOmdb string = "a24f82d6LLL"
const apiKeyOpenSub string = "mqGeQypQuxBs5bJG4pkq0Fj6Q4XKWF7bCCC"

type Movie struct {
	Title  string `json:"Title"`
	ImdbID string `json:"imdbID"`
}

func GetMovieInfo(title string) (*Movie, error) {
	// create the url
	url := fmt.Sprintf("http://www.omdbapi.com/?t=%s&apikey=%s", title, apiKeyOmdb)

	// make the GET request
	response, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}

	// close body after this function be executed
	defer response.Body.Close()

	// read the body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// unmarshal => json to struct
	var movie Movie
	err = json.Unmarshal(body, &movie)

	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response JSON: %v", err)
	}

	// return the pointer to movie and without error
	return &movie, nil
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginOpenSub(username, password string) (string, error) {
	url := "https://api.opensubtitles.com/api/v1/login"

	payload := map[string]string{
		"username": username,
		"password": password,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Api-Key", apiKeyOpenSub)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(body, &loginResponse)

	if err != nil {
		return "", err
	}

	return loginResponse.Token, nil
}
