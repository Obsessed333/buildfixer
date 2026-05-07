package main

import(
	"net/http"
	"fmt"
	"encoding/json"
	"net/url"
	"github.com/obsessed333/buildfixer/internal/models"
	"io"
)


func fetchFromPoE(accountName string, charName string) (*models.CharacterData, error) {
	endpoint := "https://pathofexile.com/character-window/get-items"

	data := url.Values{}
	data.Set("accountName", accountName)
	data.Set("character", charName)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "build-fixer")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	var result models.CharacterData
	err = json.NewDecoder(resp.Body).Decode(&result)
	return &result, err
}