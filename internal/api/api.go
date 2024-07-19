package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiKeyOmdb string = ""
const apiKeyOpenSub string = ""

type Movie struct {
	Title  string `json:"Title"`
	ImdbID string `json:"imdbID"`
}

func GetMovieInfo(title string) (*Movie, error) {
	// create the url
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", apiKeyOmdb, title)

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
	req.Header.Add("User-Agent", "easysub v0.0.1")
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

func GetSubByImdbId(userToken, imdb_id string) {
	url := fmt.Sprintf("https://api.opensubtitles.com/api/v1/subtitles?imdb_id=%s", imdb_id)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "easysub v0.0.1")
	req.Header.Add("Api-Key", apiKeyOpenSub)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
