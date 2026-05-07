package main

import(
	"net/http"
	"fmt"
	"encoding/json"
	"net/url"
	"strings"
)


func fetchFromPoE(accountName string, charName string) (*CharacterData, error) {
	endpoint := "https://pathofexile.com"
	data := url.Values{}
	data.Set("accountName", accountName)
	data.Set("character", charName)

	req, _ := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "build-fixer")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var result CharacterData
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}